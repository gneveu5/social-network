package socialnetwork

import (
	"errors"
	"fmt"
	"log"

	st "socialnetwork/pkg/structs"

	database "socialnetwork/pkg/db/sqlite"
)

func GetFollowers(userselected string, usersend int) ([]st.Follow, error) {
	var followers []st.Follow
	userselectedId, err := GetIdfromNickname(userselected)
	if err != nil {
		return nil, errors.New("utilisateur introuvable")
	}
	followstatus := GetFollowSatus(userselectedId, usersend)
	publicprivate, _ := GetPublicPrivatefromId(userselectedId)
	if followstatus != 1 && publicprivate && userselectedId != usersend {
		return nil, errors.New("vous n'etes pas autorisé")
	}

	query := `
		SELECT users.id, users.nickname, users.avatar
		FROM follows
		INNER JOIN users ON follows.follower_user_id = users.id
		WHERE follows.following_user_id = ?
		AND follows.follow_status = 1
	`
	rows, err := database.Db.Query(query, userselectedId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var follower st.Follow
		if err := rows.Scan(&follower.ID, &follower.Nickname, &follower.Avatar); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func SortFollowersForGroup(followers []st.Follow, groupId string) []st.Follow {

	var followersToReturn []st.Follow

	for _, v := range followers {
		var count int
		err := database.Db.QueryRow(`
		SELECT COUNT(*)
		FROM notif
		WHERE (notif_type = 0 AND id_two = ?)
		OR (notif_type = 1 AND target_user = ?)`,
		v.ID, v.ID).Scan(&count)
		if err != nil || count > 0 {
			continue
		}
		err = database.Db.QueryRow(`
		SELECT COUNT(*)
		FROM group_members
		WHERE user_id = ? AND group_id = ?`,
		v.ID, groupId).Scan(&count)
		if err != nil || count > 0 {
			continue
		}
		followersToReturn = append(followersToReturn, v)
	}
	return followersToReturn
}

func GetFollowing(userselected string, usersend int) ([]st.Follow, error) {
	var followers []st.Follow
	userselectedId, err := GetIdfromNickname(userselected)
	if err != nil {
		return nil, errors.New("utilisateur introuvable")
	}
	followstatus := GetFollowSatus(userselectedId, usersend)
	publicprivate, _ := GetPublicPrivatefromId(userselectedId)
	if followstatus != 1 && publicprivate && userselectedId != usersend {
		return nil, errors.New("vous n'etes pas autorisé")
	}

	query := `
		SELECT users.id, users.nickname, users.avatar
		FROM follows
		INNER JOIN users ON follows.following_user_id = users.id
		WHERE follows.follower_user_id = ?
		AND follows.follow_status = 1
	`
	rows, err := database.Db.Query(query, userselectedId)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var follower st.Follow
		if err := rows.Scan(&follower.ID, &follower.Nickname, &follower.Avatar); err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

func Follow(userselected string, usersend int) error {
	userselectedId, err := GetIdfromNickname(userselected)
	if err != nil {
		return err
	}
	_, err = database.Db.Exec("INSERT INTO follows (follower_user_id, following_user_id, follow_status) VALUES (?, ?,1)", usersend, userselectedId)
	if err != nil {
		return err
	}
	return nil
}

func UnFollow(userselected string, usersend int) error {
	userselectedId, err := GetIdfromNickname(userselected)
	if err != nil {
		return err
	}

	_, err = database.Db.Exec("DELETE FROM follows WHERE follower_user_id = ? AND following_user_id = ? AND follow_status = 1", usersend, userselectedId)
	if err != nil {
		return err
	}

	return nil
}

func FollowPrivate(userselected string, usersend int) error {
	userselectedId, err := GetIdfromNickname(userselected)
	if err != nil {
		return err
	}
	_, err = database.Db.Exec("INSERT INTO follows (follower_user_id, following_user_id, follow_status) VALUES (?, ?,0)", usersend, userselectedId)
	if err != nil {
		return err
	}
	return nil
}

func CancelFollowPrivate(userselected string, usersend int) error {
	userselectedId, err := GetIdfromNickname(userselected)
	if err != nil {
		return err
	}

	_, err = database.Db.Exec("DELETE FROM follows WHERE follower_user_id = ? AND following_user_id = ? AND follow_status = 0", usersend, userselectedId)
	if err != nil {
		return err
	}

	return nil
}

func ConfirmFollow(askingUser int, targetUser int, confirm bool) error {
	var err error
	if confirm {
		_, err = database.Db.Exec(`UPDATE follows
			SET follow_status = 1
			WHERE follower_user_id = ? AND following_user_id = ?`, askingUser, targetUser)
	} else {
		_, err = database.Db.Exec(`DELETE FROM follows
			WHERE follower_user_id = ? AND following_user_id = ?`, askingUser, targetUser)
	}

	if err != nil {
		return fmt.Errorf("ConfirmFollow : querry database\n" + err.Error())
	}

	return nil
}

func GetIdfromNickname(nickname string) (int, error) {
	var userselectedId int
	err := database.Db.QueryRow("SELECT id FROM users WHERE nickname = ?", nickname).Scan(&userselectedId)
	if err != nil {
		return 0, err
	}
	return userselectedId, nil
}

func GetPublicPrivatefromId(id int) (bool, error) {
	var publicprivate bool
	err := database.Db.QueryRow("SELECT public_private FROM users WHERE id = ?", id).Scan(&publicprivate)
	if err != nil {
		return true, err
	}
	return publicprivate, nil
}
