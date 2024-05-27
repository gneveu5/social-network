package socialnetwork

import (
	"encoding/json"
	"fmt"
	"net/http"

	h "socialnetwork/pkg/api/_http"
	tok "socialnetwork/pkg/api/jwt"
	ws "socialnetwork/pkg/api/websocket"
	req "socialnetwork/pkg/db/requests"
)

func NotificationsHandler(w http.ResponseWriter, r *http.Request) {
	client, _ := tok.GetUserIdFromToken(r)

	notifications, err := req.GetNotifications(client)
	if err != nil {
		fmt.Println("Error NotificationsHandler : getting notifications\n" + err.Error())
		h.Http400(w, "NotificationsHandler : getting notifications\n"+err.Error())
		return
	}

	notificationsJSON, err := json.Marshal(notifications)
	if err != nil {
		h.Http400(w, "NotificationsHandler : getting notifications\n"+err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(notificationsJSON)

	req.SeeAllNotifications(client)
}

type NotificationResponse struct {
	NotificationId int  `json:"notificationId"`
	Confirm        bool `json:"confirm"`
}

func NotificationResponseHandler(w http.ResponseWriter, r *http.Request) {
	var response NotificationResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		fmt.Println("NotificationResponseHandler : decoding body\n" + err.Error())
		return
	}

	targetUserId, _ := tok.GetUserIdFromToken(r)

	notification, err := req.GetVerifiedNotification(response.NotificationId, targetUserId)
	if err != nil {
		fmt.Printf("NotificationResponseHandler : can't find notification %d from user %d\n"+err.Error(),
			response.NotificationId,
			targetUserId)
		return
	}

	switch notification.NotifType {
	case 0:
		if response.Confirm {
			err = req.ConfirmGroupMembership(notification.IdTwo, notification.IdOne)
			if err == nil {
				err = ws.AddGroupChat(notification.IdTwo, notification.IdOne)
			}
		}
	case 1:
		if response.Confirm {
			err = req.ConfirmGroupMembership(targetUserId, notification.IdOne)
			if err == nil {
				err = ws.AddGroupChat(targetUserId, notification.IdOne)
			}
		}
	case 2:
		err = req.ConfirmFollow(notification.IdOne, targetUserId, response.Confirm)
		if err == nil && response.Confirm {
			var mutualFollow bool
			mutualFollow, err = req.IsFollowerOf(targetUserId, notification.IdOne)
			if mutualFollow && err == nil {
				err2 := ws.AddPrivateChat(targetUserId, notification.IdOne)
				if err2 != nil {
					fmt.Println("NotificationHandler : updating chat on new mutual follow\n" + err2.Error())
				}
				err2 = ws.AddPrivateChat(notification.IdOne, targetUserId)
				if err2 != nil {
					fmt.Println("NotificationHandler : updating chat on new mutual follow\n" + err2.Error())
				}
			}
		}
	case 3:
		if response.Confirm {
			err = req.EventAttendeesRegister_Temp(notification.IdTwo, targetUserId)
		} else {
			err = req.EventAttendeesUnregister_Temp(notification.IdTwo, targetUserId)
		}
	default:
		fmt.Printf("NotificationResponseHandler : invalid notification type %d\n", notification.NotifType)
		return
	}

	if err != nil {
		fmt.Println("NotificationResponseHandler : resolving notification\n" + err.Error())
		h.Http400(w, "NotificationResponseHandler : resolving notification\n"+err.Error())
		return
	}

	err = req.RemoveNotificationFromId(response.NotificationId)
	if err != nil {
		fmt.Println("NotificationResponseHandler : failed removing notification\n" + err.Error())
		return
	}

	h.Http200(w, "ok")
}
