package socialnetwork

import (
	database "socialnetwork/pkg/db/sqlite"
	st "socialnetwork/pkg/structs"
)

func GetUserInfo(userselected string, usersend int) ([]st.UserInfo, error) {
	var userinfo []st.UserInfo

	rows, err := database.Db.Query("SELECT id, nickname, first_name, last_name, date_of_birth, about_me, avatar, public_private, created_at FROM users WHERE nickname = ?", userselected)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user st.UserInfo
		err := rows.Scan(&user.Id, &user.Nickname, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.About_me, &user.Avatar, &user.Public_Private, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		if user.Id == usersend {
			user.FollowStatus = 3
		} else {
			user.FollowStatus = GetFollowSatus(user.Id, usersend)
		}
		if user.FollowStatus == 2 || user.FollowStatus == 0 {
			if user.Public_Private {
				user.Id = 0
				user.FirstName = ""
				user.LastName = ""
				user.DateOfBirth = ""
				user.About_me = ""
				user.CreatedAt = ""

			}
		}
		userinfo = append(userinfo, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userinfo, nil
}

func GetUserIdFromNick(userselected string) (int, error) {
	var userId int
	err := database.Db.QueryRow("SELECT id FROM users WHERE nickname = ?", userselected).Scan(&userId)
	if err != nil {
		return userId, err
	}
	return userId, nil
}

func GetAvatarAndUsernameByID(userid int) (st.AvatarAndUsername, error) {
	var user st.AvatarAndUsername
	query := "SELECT nickname, avatar FROM users WHERE id = ?"
	err := database.Db.QueryRow(query, userid).Scan(&user.Nickname, &user.Avatar)
	if err != nil {
		return user, err
	}
	return user, nil
}

func GetFollowSatus(userselected, usersend int) int {
	followstatus := 0
	err := database.Db.QueryRow("SELECT follow_status FROM follows WHERE follower_user_id = ? AND following_user_id = ?", usersend, userselected).Scan(&followstatus)
	if err != nil {
		return 2
	}
	return followstatus
}

func SwapPrivatetopublic(usersend int) error {
	_, err := database.Db.Exec("UPDATE users SET public_private = ? WHERE id = ?", 0, usersend)
	if err != nil {
		return err
	}

	return nil
}

func SwapPublictoprivate(usersend int) error {
	_, err := database.Db.Exec("UPDATE users SET public_private = ? WHERE id = ?", 1, usersend)
	if err != nil {
		return err
	}

	return nil
}
