package socialnetwork

import (
	"fmt"

	database "socialnetwork/pkg/db/sqlite"
	st "socialnetwork/pkg/structs"
)

func GetNotifications(userId int) (st.Notifications, error) {
	var res st.Notifications

	databaseNotifications, err := GetDatabaseNotifications(userId)
	if err != nil {
		return res, fmt.Errorf("GetNotifications : getting notifs from database\n" + err.Error())
	}

	for _, notification := range databaseNotifications {
		switch notification.NotifType {
		case 0:
			filledNotif, err := FillGroupMembershipRequestNotif(notification)
			if err != nil {
				return res, fmt.Errorf("GetNotifications : filling groupMembershipRequest notif\n" + err.Error())
			}
			res.GroupMembershipRequest = append(res.GroupMembershipRequest, filledNotif)
		case 1:
			filledNotif, err := FillGroupInviteNotif(notification)
			if err != nil {
				return res, fmt.Errorf("GetNotifications : filling groupInvite notif\n" + err.Error())
			}
			res.GroupInvite = append(res.GroupInvite, filledNotif)
		case 2:
			filledNotif, err := FillFollowRequestNotif(notification)
			if err != nil {
				return res, fmt.Errorf("GetNotifications : filling followRequest notif\n" + err.Error())
			}
			res.FollowRequest = append(res.FollowRequest, filledNotif)
		case 3:
			filledNotif, err := FillEventNotif(notification)
			if err != nil {
				return res, fmt.Errorf("GetNotifications : filling event notif\n" + err.Error())
			}
			res.Event = append(res.Event, filledNotif)
		}
	}

	return res, nil
}

func GetDatabaseNotifications(userId int) ([]st.Notification, error) {
	var res []st.Notification

	rows, err := database.Db.Query(`SELECT id, id_one, id_two, notif_type FROM notif WHERE target_user = ?`, userId)
	if err != nil {
		return res, fmt.Errorf("GetNotifications : query database\n" + err.Error())
	}

	for rows.Next() {
		var notification st.Notification
		err = rows.Scan(&notification.Id, &notification.IdOne, &notification.IdTwo, &notification.NotifType)
		if err != nil {
			return res, fmt.Errorf("GetNotifications : scan data\n" + err.Error())
		}
		res = append(res, notification)
	}

	return res, nil
}

func FillGroupMembershipRequestNotif(notification st.Notification) (st.GroupMembershipRequestNotif, error) {
	res := st.GroupMembershipRequestNotif{Notification: notification}

	groupName, err := GetGroupNameFromId(notification.IdOne)
	if err != nil {
		return res, fmt.Errorf("FillGroupInviteNotif : getting group name\n" + err.Error())
	}

	userName, err := GetUsernameFromId(notification.IdTwo)
	if err != nil {
		return res, fmt.Errorf("FillGroupInviteNotif : getting user name\n" + err.Error())
	}

	res.GroupName = groupName
	res.UserName = userName

	return res, nil
}

func FillGroupInviteNotif(notification st.Notification) (st.GroupInviteNotif, error) {
	var err error

	res := st.GroupInviteNotif{Notification: notification}

	groupName, err := GetGroupNameFromId(notification.IdOne)
	if err != nil {
		return res, fmt.Errorf("FillGroupInviteNotif : getting group name\n" + err.Error())
	}

	userName, err := GetUsernameFromId(notification.IdTwo)
	if err != nil {
		return res, fmt.Errorf("FillGroupInviteNotif : getting group name\n" + err.Error())
	}

	res.GroupName = groupName
	res.UserName = userName

	return res, nil
}

func FillFollowRequestNotif(notification st.Notification) (st.FollowRequestNotif, error) {
	res := st.FollowRequestNotif{Notification: notification}

	userName, err := GetUsernameFromId(notification.IdOne)
	if err != nil {
		return res, fmt.Errorf("FillFollowRequestNotif : getting username\n" + err.Error())
	}

	res.UserName = userName

	return res, nil
}

func FillEventNotif(notification st.Notification) (st.EventNotif, error) {
	res := st.EventNotif{Notification: notification}

	groupName, err := GetGroupNameFromId(notification.IdOne)
	if err != nil {
		return res, fmt.Errorf("FillEventNotif : getting group name\n" + err.Error())
	}

	eventName, err := GetEventNameFromId(notification.IdTwo)
	if err != nil {
		return res, fmt.Errorf("FillEventNotif : getting event name\n" + err.Error())
	}

	res.GroupName = groupName
	res.EventName = eventName

	return res, nil
}

func GetGroupNameFromId(groupId int) (string, error) {
	var res string

	err := database.Db.QueryRow(`SELECT title FROM users_group WHERE id = ?`, groupId).Scan(&res)
	if err != nil {
		return "", err
	}

	return res, nil
}

func GetEventNameFromId(eventId int) (string, error) {
	var res string

	err := database.Db.QueryRow(`SELECT title FROM group_event WHERE id = ?`, eventId).Scan(&res)
	if err != nil {
		return "", err
	}

	return res, nil
}

func RemoveNotificationFromId(notifId int) error {
	_, err := database.Db.Exec(`DELETE FROM notif WHERE id = ?`, notifId)

	return err
}

func GetVerifiedNotification(notifId int, userId int) (st.Notification, error) {
	var res st.Notification

	err := database.Db.QueryRow(`
		SELECT id, id_one, id_two, notif_type FROM notif WHERE id = ? AND target_user = ?`, notifId, userId).Scan(
		&res.Id, &res.IdOne, &res.IdTwo, &res.NotifType)

	return res, err
}

func RemoveNotificationFromDatas(notifType int, targetUser int, idOne int, idTwo int) error {
	_, err := database.Db.Exec(`DELETE FROM notif WHERE
		notif_type = ? AND target_user = ? AND id_one = ? and id_two = ?`,
		notifType, targetUser, idOne, idTwo)

	return err
}

func GetNotificationsCount(userId int) (int, error) {
	var count int
	err := database.Db.QueryRow(`SELECT count(*) FROM notif WHERE target_user = ? AND seen = 0`, userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("GetNotificationsCount : database req\n" + err.Error())
	}

	return count, nil
}

func SeeAllNotifications(userId int) error {
	_, err := database.Db.Exec(`UPDATE notif
		SET seen = 1
		WHERE target_user = ? AND seen = 0
		`, userId)

	return err
}

func InsertNotif(targetUser int, notifType int, idOne int, idTwo int) error {
	_, err := database.Db.Exec(`INSERT INTO notif
		(target_user, id_one, id_two, notif_type)
		VALUES(?,?,?,?)`, targetUser, idOne, idTwo, notifType)

	return err
}
