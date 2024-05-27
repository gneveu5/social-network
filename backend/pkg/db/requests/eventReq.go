package socialnetwork

import (
	"database/sql"
	"fmt"

	database "socialnetwork/pkg/db/sqlite"
	model "socialnetwork/pkg/structs"

	"github.com/golang-jwt/jwt/v5"
)

func EventInserter(event model.EventReturn, userId int) ([]model.GroupMembers, error) {
	_, err := database.Db.Exec("INSERT INTO group_event (title, event_description, event_time, user_id, users_group_id) VALUES (?, ?, ?, ?, ?)", event.Title, event.Description, event.EventDate, userId, event.GroupId)
	if err != nil {
		return nil, err
	}

	var eventId int
	err = database.Db.QueryRow("SELECT last_insert_rowid() FROM group_event").Scan(&eventId)
	if err != nil {
		return nil, err
	}

	groupMembers, err := GroupMembersFetcher(userId, event.GroupId)
	if err != nil {
		return nil, err
	}

	for _, v := range groupMembers {
		_, err := database.Db.Exec("INSERT INTO notif (target_user, id_one, id_two, notif_type) VALUES (?, ?, ?, 3)", v.UserId, event.GroupId, eventId)
		if err != nil {
			return nil, err
		}
	}

	return groupMembers, nil
}

func EventFetcher(groupId string) ([]model.Event, error) {
	var rows *sql.Rows
	var event model.Event
	var err error
	var eventList []model.Event

	rows, err = database.Db.Query(`
	SELECT DISTINCT group_event.*, users.nickname
	FROM group_event
	LEFT JOIN users ON users.id = group_event.user_id
	WHERE group_event.users_group_id = ?
	ORDER BY group_event.id DESC
	`, groupId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&event.Id,
			&event.Title,
			&event.Description,
			&event.EventDate,
			&event.CreatedAt,
			&event.UserId,
			&event.GroupId,
			&event.CreatorName,
		)
		if err != nil {
			return nil, err
		}
		eventList = append(eventList, event)
	}
	return eventList, nil
}

func EventAttendeesFetcher(eventId string) ([]model.EventAttendees, error) {
	var rows *sql.Rows
	var userEvent model.EventAttendees
	var err error
	var userEventList []model.EventAttendees

	rows, err = database.Db.Query(`
	SELECT DISTINCT user_event.*, users.nickname
	FROM user_event
	LEFT JOIN users ON users.id = user_event.user_id
	WHERE user_event.event_id = ?
	ORDER BY users.nickname
	`, eventId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&userEvent.Id,
			&userEvent.UserId,
			&userEvent.EventId,
			&userEvent.AttendeeName,
		)
		if err != nil {
			return nil, err
		}
		userEventList = append(userEventList, userEvent)
	}
	return userEventList, nil
}

func EventAttendeesRegister(eventId int, claims jwt.MapClaims) error {
	_, err := database.Db.Exec("INSERT INTO user_event (user_id, event_id) VALUES (?, ?)", claims["id"], eventId)
	if err != nil {
		return err
	}
	return nil
}

func EventAttendeesUnregister(eventId int, claims jwt.MapClaims) error {
	_, err := database.Db.Exec("DELETE FROM user_event WHERE user_id=? AND event_id=?", claims["id"], eventId)
	if err != nil {
		return err
	}
	return nil
}

func EventAttendeesStatus(eventId int, claims jwt.MapClaims) (int, error) {
	var row int
	err := database.Db.QueryRow(`
	SELECT DISTINCT user_event.event_id
	FROM user_event
	WHERE user_event.event_id = ?
	AND user_event.user_id = ?
	`, eventId, claims["id"]).Scan(&row)
	return row, err
}

// à changer un jour
func EventAttendeesUnregister_Temp(eventId int, userId int) error {
	_, err := database.Db.Exec("DELETE FROM user_event WHERE user_id=? AND event_id=?", userId, eventId)

	return err
}

func EventAttendeesRegister_Temp(eventId int, userId int) error {
	err := EventAttendeesUnregister_Temp(eventId, userId)
	if err != nil {
		return fmt.Errorf("EventAttendedRegister : trying to remove previous registrations\n" + err.Error())
	}

	_, err = database.Db.Exec("INSERT INTO user_event (user_id, event_id) VALUES (?, ?)", userId, eventId)

	return err
}
