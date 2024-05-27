package socialnetwork

import (
	st "socialnetwork/pkg/structs"

	database "socialnetwork/pkg/db/sqlite"
)

func GetUserNickname(search string) ([]st.UserSearch, error) {
	if search == "" || string(search[0]) == " " {
		return nil, nil
	}
	query := "SELECT id, nickname FROM users WHERE nickname LIKE '%" + search + "%' LIMIT 10"
	rows, err := database.Db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []st.UserSearch
	count := 0

	for rows.Next() {
		if count >= 10 {
			break
		}
		var id int
		var nickname string
		if err := rows.Scan(&id, &nickname); err != nil {
			return nil, err
		}
		user := st.UserSearch{
			Id:       id,
			Nickname: nickname,
		}
		users = append(users, user)
		count++
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
