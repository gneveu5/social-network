package socialnetwork

import (
	"fmt"
	"sync"

	conf "socialnetwork/config"
	req "socialnetwork/pkg/db/requests"
	st "socialnetwork/pkg/structs"

	"github.com/gorilla/websocket"
)

// stores all datas related to users connections
var (
	ConnectedUsers      map[int]st.ConnectedUserData = make(map[int]st.ConnectedUserData)
	ConnectedUsersMutex sync.Mutex
)

// build user datas from userid
func GetUserData(userId int) (st.ConnectedUserData, error) {
	var res st.ConnectedUserData

	name, err := req.GetUsernameFromId(userId)
	if err != nil {
		return res, fmt.Errorf("GetUserData : requesting username\n" + err.Error())
	}

	res.Name = name

	return res, nil
}

// get all groups that users belongs to
func GetGroupsList(userId int) ([]st.Contact, error) {
	res, err := req.GetGroupsFromUserId(userId)
	if err != nil {
		return res, fmt.Errorf("GetGroupsList : getting groups id\n" + err.Error())
	}

	for i, group := range res {
		res[i].History, err = req.GetGroupMessagesHistory(group.ContactId, 0)
		if err != nil {
			return res, fmt.Errorf("GetGroupsList : getting history\n" + err.Error())
		}
	}

	return res, nil
}

// return list of all user's contacts from database
func GetPrivateContacts(userId int) ([]st.PrivateContact, error) {
	res, err := req.GetContactsFromUserId(userId)
	if err != nil {
		return res, fmt.Errorf("GetAllContacts : getting contacts\n" + err.Error())
	}

	for i, c := range res {
		res[i].History, err = req.GetPrivateMessagesHistory(userId, c.ContactId, 0)
		if err != nil {
			return res, fmt.Errorf("GetAllContacts : getting history\n" + err.Error())
		}
	}

	return res, nil
}

// handles websocket's termination
func DisconnectUser(userId int, conn *websocket.Conn) error {
	// update ConnectedUsers. delete entry if no more websocket
	ConnectedUsersMutex.Lock()
	var updatedConns []*websocket.Conn
	for _, c := range ConnectedUsers[userId].Conns {
		if c != conn {
			updatedConns = append(updatedConns, c)
		}
	}
	if len(updatedConns) == 0 {
		// no more device connected
		delete(ConnectedUsers, userId)
	} else {
		// some devices are still connected
		data := ConnectedUsers[userId]
		data.Conns = updatedConns
		ConnectedUsers[userId] = data
	}
	ConnectedUsersMutex.Unlock()

	// send disconnect message to all contacts
	var disconnectMsg st.UpdateConnectedStatus
	disconnectMsg.UserId = userId

	wsm, err := BuildWebsocketMessage("disconnect", disconnectMsg)
	if err != nil {
		return fmt.Errorf("DisconnectUser : building disconnect wsm\n" + err.Error())
	}

	contacts, _ := GetPrivateContacts(userId)
	for _, contact := range contacts {
		SendToUser(wsm, contact.ContactId)
	}

	return nil
}

// handles websocket opening
func ConnectUser(userId int, conn *websocket.Conn) error {
	var data st.ConnectedUserData

	privates, err := GetPrivateContacts(userId)
	if err != nil {
		return err
	}

	groups, err := GetGroupsList(userId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// init messages to send through websocket
	connectMsg := st.UpdateConnectedStatus{UserId: userId}
	wsmConnect, err := BuildWebsocketMessage("connect", connectMsg)
	if err != nil {
		return fmt.Errorf("ConnectUser : building connect wsm\n" + err.Error())
	}

	contactListMsg := st.ContactListMsg{Group: groups}

	ConnectedUsersMutex.Lock()

	// creates new entry if user is not logged in, updates websockets list otherwise
	_, alreadyConnected := ConnectedUsers[userId]
	if alreadyConnected {
		data = ConnectedUsers[userId]
	} else {
		data, err = GetUserData(userId)
		if err != nil {
			ConnectedUsersMutex.Unlock()
			return err
		}
	}
	data.Conns = append(data.Conns, conn)

	// close oldest websocket if too much are opened
	if len(data.Conns) > conf.MAX_WEBSOCKETS {
		data.Conns[0].Close()
		data.Conns = data.Conns[1:]
	}

	ConnectedUsers[userId] = data

	// iterate through contact list
	// update connected status in contact list and send connect message to online contacts
	for i, privateContact := range privates {
		if _, connected := ConnectedUsers[privateContact.ContactId]; connected {
			for _, c := range ConnectedUsers[privateContact.ContactId].Conns {
				// c.WriteJSON(connectMsg)
				c.WriteJSON(wsmConnect)
			}
			privates[i].Connected = true
		} else {
			privates[i].Connected = false
		}
	}

	ConnectedUsersMutex.Unlock()

	// send contact list to user
	contactListMsg.Private = privates

	wsmContactList, err := BuildWebsocketMessage("init-contact", contactListMsg)
	if err != nil {
		return fmt.Errorf("ConnectUser : building contactList wsm\n" + err.Error())
	}
	conn.WriteJSON(wsmContactList)

	return nil
}
