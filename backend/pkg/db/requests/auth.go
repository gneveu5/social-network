package socialnetwork

import (
	st "socialnetwork/pkg/structs"

	database "socialnetwork/pkg/db/sqlite"
)

func GetUserByEmail(email string) (st.User, error) {
	var user st.User
	query := "SELECT id, nickname, email, user_password, first_name, last_name, date_of_birth, about_me, public_private, avatar, created_at, has_validated FROM users WHERE email = ?"
	err := database.Db.QueryRow(query, email).Scan(&user.ID, &user.NickName, &user.Email, &user.Password, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.AboutMe, &user.PublicPrivate, &user.Avatar, &user.CreatedAt, &user.HasValidated)
	if err != nil {
		return user, err
	}
	return user, nil
}

func InsertUser(user st.SingUpRequest) error {
	_, err := database.Db.Exec("INSERT INTO users (email, user_password, first_name, last_name, date_of_birth, avatar, nickname, about_me, public_private, has_validated) VALUES (?, ?, ?, ?, ?, ?, ?, ?, 0, 1)", user.Email, user.Password, user.FirstName, user.LastName, user.DateOfBirth, user.AvatarName, user.NickName, user.AboutMe)
	if err != nil {
		return err
	}
	return nil
}

func CheckEmail(email string) bool {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email).Scan(&count)
	if err != nil {
		return false
	}
	return count == 0
}

func CheckNickname(nickname string) bool {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM users WHERE nickname = ?", nickname).Scan(&count)
	if err != nil {
		return false
	}
	return count == 0
}
