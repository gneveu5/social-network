package socialnetwork

import (
	"fmt"

	req "socialnetwork/pkg/db/requests"
	st "socialnetwork/pkg/structs"

	"github.com/gorilla/websocket"
)

func SendNotificationsCount(userId int, conn *websocket.Conn) {
	count, err := req.GetNotificationsCount(userId)
	if err != nil {
		fmt.Println("SendNotificationsCount : getting count\n" + err.Error())
		return
	}

	if count == 0 {
		return
	}

	data := st.NotificationCount{Count: count}

	wsm, err := BuildWebsocketMessage("notif", data)
	if err != nil {
		fmt.Println("SendNotificationsCount : building websocket message\n" + err.Error())
		return
	}

	conn.WriteJSON(wsm)
}

func SendNewNotification(userId int) error {
	wsm, err := BuildWebsocketMessage("newnotif", nil)
	if err != nil {
		return fmt.Errorf("SendNewNotification : building wsm\n" + err.Error())
	}

	SendToUser(wsm, userId)

	return nil
}
