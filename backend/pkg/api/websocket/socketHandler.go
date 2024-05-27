package socialnetwork

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	conf "socialnetwork/config"
	tok "socialnetwork/pkg/api/jwt"
	st "socialnetwork/pkg/structs"

	"github.com/gorilla/websocket"
)

// handles socket connection
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
		// ReadBufferSize:  1024,
		// WriteBufferSize: 1024,
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("failed to upgrade conn")
		fmt.Println(err)
		return
	}

	defer conn.Close()

	_, jwt, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Socket handler : failed reading jwt\n" + err.Error())
		return
	}

	claims, err := tok.VerifyToken(string(jwt))
	if err != nil {
		fmt.Println("failed retrieveing jwt in websocket access")
		fmt.Println(err)
		return
	}
	userIdStr := fmt.Sprintf("%v", claims["id"])

	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		fmt.Println("invalid userId in websocket access : " + userIdStr)
		return
	}

	err = ConnectUser(userId, conn)
	if err != nil {
		fmt.Println("Socket handler : connecting user\n" + err.Error())
		return
	}

	SendNotificationsCount(userId, conn)

	var lastMessageTimeStamp int64 = time.Now().UnixMilli() - conf.REQUEST_SPAM_DELAY

	for {
		var wsMsg st.WebsocketMessage
		err := conn.ReadJSON(&wsMsg)
		if err != nil {
			DisconnectUser(userId, conn)
			return
		}

		currentTimeStamp := time.Now().UnixMilli()
		if currentTimeStamp-lastMessageTimeStamp < conf.REQUEST_SPAM_DELAY {
			continue
		} else {
			lastMessageTimeStamp = currentTimeStamp
		}

		switch wsMsg.Type {
		case "pvt-msg":
			var msg st.ChatMessage
			err = json.Unmarshal([]byte(wsMsg.Data), &msg)
			if err != nil {
				err = fmt.Errorf("SocketHandler : unmarshal pvt-msg\n" + err.Error())
			} else {
				err = PrivateMessage(msg, userId)
			}
		case "grp-msg":
			var msg st.ChatMessage
			err = json.Unmarshal([]byte(wsMsg.Data), &msg)
			if err != nil {
				err = fmt.Errorf("SocketHandler : unmarshal grp-msg\n" + err.Error())
			} else {
				err = GroupMessage(msg, userId)
			}
		case "pvt-hist":
			var msg st.HistoryRequest
			err = json.Unmarshal([]byte(wsMsg.Data), &msg)
			if err != nil {
				err = fmt.Errorf("SocketHandler : unmarshal pvt-hist\n" + err.Error())
			} else {
				err = SendPrivateHistory(userId, msg, conn)
			}
		case "grp-hist":
			var msg st.HistoryRequest
			err = json.Unmarshal([]byte(wsMsg.Data), &msg)
			if err != nil {
				err = fmt.Errorf("SocketHandler : unmarshal grp-hist\n" + err.Error())
			} else {
				err = SendGroupHistory(msg, conn)
			}
		default:
			fmt.Printf("message type %s not handled\n", wsMsg.Type)
			fmt.Println(wsMsg.Data)
		}
		if err != nil {
			fmt.Println("Socket handler :\n" + err.Error())
		}
	}
}
