package socialnetwork

import (
	"database/sql"
	"fmt"
	"slices"
	"time"

	conf "socialnetwork/config"
	database "socialnetwork/pkg/db/sqlite"
	st "socialnetwork/pkg/structs"
)

// returns all user's contacts (i.e mutual follow between user and contact)
func GetContactsFromUserId(userId int) ([]st.PrivateContact, error) {
	var res []st.PrivateContact

	rows, err := database.Db.Query(`
		SELECT id, nickname FROM users WHERE id IN (
			SELECT A.follower_user_id FROM follows A, follows B WHERE
			A.follow_status = 1
			AND B.follow_status = 1
			AND A.following_user_id = ? 
			AND B.follower_user_id = ?
			AND A.follower_user_id = B.following_user_id 
		)
	`, userId, userId)
	if err != nil {
		return res, fmt.Errorf("GetContactsFromUserId : query database\n" + err.Error())
	}

	for rows.Next() {
		var contactId int
		var name string
		err = rows.Scan(&contactId, &name)
		if err != nil {
			return res, fmt.Errorf("GetContactsFromUserId : scan data\n" + err.Error())
		}

		res = append(res, st.PrivateContact{Contact: st.Contact{ContactId: contactId, Name: name}})
	}

	return res, nil
}

// returns all group that user belongs to
func GetGroupsFromUserId(userId int) ([]st.Contact, error) {
	var res []st.Contact

	rows, err := database.Db.Query(`
		SELECT id, title FROM users_group WHERE id IN (
			SELECT group_id FROM group_members WHERE user_id = ?
		)
	`, userId)
	if err != nil {
		return res, fmt.Errorf("GetGroupsFromUserId : query database\n" + err.Error())
	}

	for rows.Next() {
		var groupId int
		var name string
		err = rows.Scan(&groupId, &name)
		if err != nil {
			return res, fmt.Errorf("GetGroupsFromUserId : scan data\n" + err.Error())
		}

		res = append(res, st.Contact{ContactId: groupId, Name: name})
	}

	return res, nil
}

func GetGroupContact(groupId int) (st.Contact, error) {
	var res st.Contact

	err := database.Db.QueryRow(`
		SELECT id, title FROM users_group WHERE id = ?`, groupId).Scan(&res.ContactId, &res.Name)
	if err != nil {
		return res, fmt.Errorf("GetGroupContactFromId : querry database\n" + err.Error())
	}

	res.History, err = GetGroupMessagesHistory(groupId, 0)
	if err != nil {
		return res, fmt.Errorf("GetGroupContactFromId : getting message history\n" + err.Error())
	}

	return res, nil
}

func GetPrivateContact(contactId int, targetId int) (st.PrivateContact, error) {
	contact := st.Contact{ContactId: contactId}

	var res st.PrivateContact

	userName, err := GetUsernameFromId(contactId)
	if err != nil {
		return res, fmt.Errorf("GetPrivateContact : getting user name\n" + err.Error())
	}

	history, err := GetPrivateMessagesHistory(contactId, targetId, 0)
	if err != nil {
		return res, fmt.Errorf("GetPrivateContact : getting message history\n" + err.Error())
	}

	contact.Name = userName
	contact.History = history

	res.Contact = contact

	return res, nil
}

// returns all group members' ids
func GetGroupMembersId(groupId int) ([]int, error) {
	var res []int

	rows, err := database.Db.Query(`
		SELECT user_id FROM group_members WHERE group_id = ?
	`, groupId)
	if err != nil {
		return res, fmt.Errorf("GetGroupMembersId : query database\n" + err.Error())
	}

	for rows.Next() {
		var userId int
		err = rows.Scan(&userId)
		if err != nil {
			return res, fmt.Errorf("GetGroupMembersId : scan data\n" + err.Error())
		}

		res = append(res, userId)
	}

	return res, nil
}

// insert message into database. type 0 for private message, 1 for group message
func InsertMessage(senderId int, message st.ChatMessage, msgType int) (int, error) {
	res, err := database.Db.Exec(`
		INSERT INTO chat_messages (author_id, target_id, user_message, message_type) VALUES (?,?,?,?)
	`, senderId, message.To, message.Txt, msgType)
	if err != nil {
		return 0, fmt.Errorf("InsertMessage : inserting into db\n" + err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("InsertMessage : retrieveing msg id\n" + err.Error())
	}

	return int(id), nil
}

func GetGroupMessagesHistory(groupId int, previousMsgId int) ([]st.ChatHistoryElement, error) {
	var (
		res  []st.ChatHistoryElement
		rows *sql.Rows
		err  error
	)

	if previousMsgId == 0 {
		rows, err = database.Db.Query(`
		SELECT cm.id, us.nickname, cm.created_at, cm.user_message FROM
		chat_messages cm INNER JOIN users us ON cm.author_id = us.id
		WHERE cm.target_id = ? AND cm.message_type = 1
		ORDER BY cm.id DESC LIMIT ?
	`, groupId, conf.HISTORY_MESSAGES_COUNT)
	} else {
		rows, err = database.Db.Query(`
		SELECT cm.id, us.nickname, cm.created_at, cm.user_message FROM
		chat_messages cm INNER JOIN users us ON cm.author_id = us.id
		WHERE cm.target_id = ? AND cm.message_type = 1 AND cm.id < ?
		ORDER BY cm.id DESC LIMIT ?
	`, groupId, previousMsgId, conf.HISTORY_MESSAGES_COUNT)
	}

	if err != nil {
		return res, fmt.Errorf("GetGroupMessageHistory : query database\n" + err.Error())
	}

	return ScanChatHistoryRequestRows(rows)
}

func GetPrivateMessagesHistory(user1 int, user2 int, previousMsgId int) ([]st.ChatHistoryElement, error) {
	var (
		res  []st.ChatHistoryElement
		rows *sql.Rows
		err  error
	)

	if previousMsgId == 0 {
		rows, err = database.Db.Query(`
		SELECT cm.id, us.nickname, cm.created_at, cm.user_message FROM
		chat_messages cm INNER JOIN users us ON cm.author_id = us.id
		WHERE cm.message_type = 0
		AND ((cm.author_id = ? AND cm.target_id = ?) OR (cm.author_id = ? AND cm.target_id = ?))
		ORDER BY cm.id DESC LIMIT ?
	`, user1, user2, user2, user1, conf.HISTORY_MESSAGES_COUNT)
	} else {
		rows, err = database.Db.Query(`
		SELECT cm.id, us.nickname, cm.created_at, cm.user_message FROM
		chat_messages cm INNER JOIN users us ON cm.author_id = us.id
		WHERE cm.message_type = 0
		AND ((cm.author_id = ? AND cm.target_id = ?) OR (cm.author_id = ? AND cm.target_id = ?))
		AND cm.id < ?
		ORDER BY cm.id DESC LIMIT ?
	`, user1, user2, user2, user1, previousMsgId, conf.HISTORY_MESSAGES_COUNT)
	}

	if err != nil {
		return res, fmt.Errorf("GetPrivateMessageHistory : query database\n" + err.Error())
	}

	return ScanChatHistoryRequestRows(rows)
}

func ScanChatHistoryRequestRows(rows *sql.Rows) ([]st.ChatHistoryElement, error) {
	var res []st.ChatHistoryElement

	for rows.Next() {
		var (
			id   int
			from string
			date time.Time
			txt  string
		)

		err := rows.Scan(&id, &from, &date, &txt)
		if err != nil {
			return res, fmt.Errorf("ScanChatHistoryRequestRows\n" + err.Error())
		}

		res = append(res, st.ChatHistoryElement{Id: id, From: from, Date: date.Unix(), Txt: txt})
	}

	slices.Reverse(res)

	return res, nil
}

func GetUsernameFromId(userId int) (string, error) {
	var name string

	err := database.Db.QueryRow("SELECT nickname FROM users WHERE id=?", userId).Scan(&name)
	if err != nil {
		return "", fmt.Errorf("GetUserDataFromUserId :\n" + err.Error())
	}

	return name, nil
}

// return true if user1 follows user2
func IsFollowerOf(user1 int, user2 int) (bool, error) {
	var count int

	err := database.Db.QueryRow(`SELECT count(*) FROM follows WHERE 
		follower_user_id = ? AND following_user_id = ?
	`, user1, user2).Scan(&count)
	if err != nil {
		return false, fmt.Errorf("IsFollowerOf : querrying database\n" + err.Error())
	}

	return count != 0, nil
}
