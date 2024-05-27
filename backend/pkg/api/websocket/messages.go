package socialnetwork

import (
	"encoding/json"
	"fmt"
	"time"

	req "socialnetwork/pkg/db/requests"
	st "socialnetwork/pkg/structs"

	"github.com/gorilla/websocket"
)

// build a websocket message from type and data
func BuildWebsocketMessage(msgType string, data interface{}) (st.WebsocketMessage, error) {
	jsonStr, err := json.Marshal(data)
	if err != nil {
		return st.WebsocketMessage{Type: "", Data: ""}, err
	}

	return st.WebsocketMessage{Type: msgType, Data: string(jsonStr)}, nil
}

// send websocketmessage to all connected devices of the user
func SendToUser(wsm st.WebsocketMessage, userId int) {
	ConnectedUsersMutex.Lock()
	userData, connected := ConnectedUsers[userId]
	if connected {
		for _, conn := range userData.Conns {
			conn.WriteJSON(wsm)
		}
	}
	ConnectedUsersMutex.Unlock()
}

// handles receiving private message data from client
// inserts message into database, sends message to recipient and sends back to client
func PrivateMessage(message st.ChatMessage, senderId int) error {
	id, err := req.InsertMessage(senderId, message, 0)
	if err != nil {
		return fmt.Errorf("PrivateMessage : inserting into database\n" + err.Error())
	}

	senderName, err := req.GetUsernameFromId(senderId)
	if err != nil {
		return fmt.Errorf("PrivateMessage : getting username from id\n" + err.Error())
	}

	responseToClients := st.ChatMessageToClient{
		ChatMessage: message,
		Id:          id,
		From:        senderName,
		Date:        time.Now().Unix(),
	}

	// send message back to sender
	wsm, err := BuildWebsocketMessage("pvt-msg", responseToClients)
	if err != nil {
		return fmt.Errorf("PrivateMessage : building pvt-msg\n" + err.Error())
	}
	SendToUser(wsm, senderId)

	// send message to recipient
	recipientId := message.To
	responseToClients.To = senderId
	wsm, err = BuildWebsocketMessage("pvt-msg", responseToClients)
	if err != nil {
		return fmt.Errorf("PrivateMessage : building pvt-msg\n" + err.Error())
	}
	SendToUser(wsm, recipientId)

	return nil
}

// handles receiving group message data from client
// inserts message into database and sends message to all group members (including client)
func GroupMessage(message st.ChatMessage, senderId int) error {
	id, err := req.InsertMessage(senderId, message, 1)
	if err != nil {
		return fmt.Errorf("GroupMessage : inserting into database\n" + err.Error())
	}

	senderName, err := req.GetUsernameFromId(senderId)
	if err != nil {
		return fmt.Errorf("PrivateMessage : getting username from id\n" + err.Error())
	}

	groupMembersId, err := req.GetGroupMembersId(message.To)
	if err != nil {
		return fmt.Errorf("GroupMessage : getting group list\n" + err.Error())
	}

	responseToClients := st.ChatMessageToClient{
		ChatMessage: message,
		Id:          id,
		From:        senderName,
		Date:        time.Now().Unix(),
	}

	wsm, err := BuildWebsocketMessage("grp-msg", responseToClients)
	if err != nil {
		return fmt.Errorf("GroupMessage : building wsm\n" + err.Error())
	}

	for _, recipientId := range groupMembersId {
		SendToUser(wsm, recipientId)
	}

	return nil
}

func SendGroupHistory(request st.HistoryRequest, conn *websocket.Conn) error {
	hist, err := req.GetGroupMessagesHistory(request.Id, request.Last)
	if err != nil {
		return fmt.Errorf("SendGroupHistory : getting history\n" + err.Error())
	}

	if len(hist) == 0 {
		return nil
	}

	var data st.HistoryResponse
	data.To = request.Id
	data.Hist = hist

	wsm, err := BuildWebsocketMessage("grp-hist", data)
	if err != nil {
		return fmt.Errorf("SendGroupHistory : building wsm\n" + err.Error())
	}

	conn.WriteJSON(wsm)

	return nil
}

func SendPrivateHistory(userId int, request st.HistoryRequest, conn *websocket.Conn) error {
	history, err := req.GetPrivateMessagesHistory(userId, request.Id, request.Last)
	if err != nil {
		return fmt.Errorf("SendGroupHistory : getting history\n" + err.Error())
	}

	if len(history) == 0 {
		return nil
	}

	var data st.HistoryResponse
	data.To = request.Id
	data.Hist = history

	wsm, err := BuildWebsocketMessage("pvt-hist", data)
	if err != nil {
		return fmt.Errorf("SendPrivateHistory : building wsm\n" + err.Error())
	}

	conn.WriteJSON(wsm)

	return nil
}

func AddGroupChat(userId int, groupId int) error {
	contact, err := req.GetGroupContact(groupId)
	if err != nil {
		return fmt.Errorf("AddGroupChat : getting contact\n" + err.Error())
	}

	contactListMsg := st.ContactListMsg{Group: []st.Contact{contact}}
	wsm, err := BuildWebsocketMessage("init-contact", contactListMsg)
	if err != nil {
		return fmt.Errorf("AddGroupChat : building wsm\n" + err.Error())
	}

	SendToUser(wsm, userId)

	return nil
}

func RemoveGroupChat(userId int, groupId int) error {
	return nil
}

func AddPrivateChat(newContactId int, targetId int) error {
	contact, err := req.GetPrivateContact(newContactId, targetId)
	if err != nil {
		return fmt.Errorf("AddPrivateChat : getting contact\n" + err.Error())
	}

	_, connected := ConnectedUsers[newContactId]
	contact.Connected = connected

	contactListMsg := st.ContactListMsg{Private: []st.PrivateContact{contact}}
	wsm, err := BuildWebsocketMessage("init-contact", contactListMsg)
	if err != nil {
		return fmt.Errorf("AddPrivateChat : building wsm\n" + err.Error())
	}

	SendToUser(wsm, targetId)

	return nil
}

func RemovePrivateChat(user1 int, user2 int) error {
	removeContact := st.RemoveContact{Id: user1}
	wsm, err := BuildWebsocketMessage("rmv-pvt", removeContact)
	if err != nil {
		return fmt.Errorf("RemovePrivateChat : building first wsm\n" + err.Error())
	}

	SendToUser(wsm, user2)

	removeContact.Id = user2
	wsm, err = BuildWebsocketMessage("rmv-pvt", removeContact)
	if err != nil {
		return fmt.Errorf("RemovePrivateChat : building second wsm\n" + err.Error())
	}

	SendToUser(wsm, user1)

	return nil
}
