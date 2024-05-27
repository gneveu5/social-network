package socialnetwork

import (
	database "socialnetwork/pkg/db/sqlite"
	model "socialnetwork/pkg/structs"

	"github.com/golang-jwt/jwt/v5"
)

func GroupCreation(groupRegistrationReturn model.GroupCreationReturn, userId interface{}) (int, error) {
	_, err := database.Db.Exec("INSERT INTO users_group (title, group_description, user_id) VALUES (?, ?, ?)", groupRegistrationReturn.Title, groupRegistrationReturn.Description, userId)
	if err != nil {
		return 0, err
	}

	var groupId int
	err = database.Db.QueryRow("SELECT last_insert_rowid() FROM users_group").Scan(&groupId)
	if err != nil {
		return 0, err
	}

	_, err = database.Db.Exec("INSERT INTO group_members (user_id, group_id) VALUES (?, ?)", userId, groupId)
	if err != nil {
		return 0, err
	}

	return groupId, nil
}

func GroupList(userId string) ([]model.Group, error) {
	var group model.Group
	var groupList []model.Group

	rows, err := database.Db.Query(`
	SELECT DISTINCT users_group.*
	FROM group_members
	LEFT JOIN users ON users.id = group_members.user_id
	LEFT JOIN users_group ON users_group.id = group_members.group_id
	WHERE users.nickname = ?
	ORDER BY users_group.id DESC
	`, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&group.Id,
			&group.Title,
			&group.Description,
			&group.UserId,
			&group.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		group.AdminName, err = GroupAdminNameFetcher(group.Id)
		if err != nil {
			return nil, err
		}
		groupList = append(groupList, group)
	}
	return groupList, nil
}

func GroupFetchAskRegister(userId interface{}, groupId string) (bool, error) {
	var count int
	err := database.Db.QueryRow("SELECT COUNT(*) FROM notif WHERE id_one = ? AND id_two = ?", groupId, userId).Scan(&count)
	if err != nil {
		return false, err
	}
	return count != 0, err
}

func GroupFetcher(groupId string) (model.Group, error) {
	var group model.Group
	err := database.Db.QueryRow("SELECT users_group.title, users_group.id FROM users_group WHERE id = ?", groupId).Scan(&group.Title, &group.Id)
	return group, err
}

func GroupAdminFetcher(groupId string) (int, error) {
	var userId int
	err := database.Db.QueryRow("SELECT users_group.user_id FROM users_group WHERE id = ?", groupId).Scan(&userId)
	return userId, err
}

func GroupAdminNameFetcher(adminId int) (string, error) {
	var userName string
	err := database.Db.QueryRow("SELECT users.nickname FROM users WHERE id = ?", adminId).Scan(&userName)
	return userName, err
}

func GroupAskRegister(userId interface{}, groupRegistrationReturn model.GroupAskMembershipReturn) error {
	var groupAdminId int
	err := database.Db.QueryRow("SELECT user_id FROM users_group WHERE id = ?", groupRegistrationReturn.GroupId).Scan(&groupAdminId)
	if err != nil {
		return err
	}

	_, err = database.Db.Exec("INSERT INTO notif (target_user, id_one, id_two, notif_type) VALUES (?, ?, ?, 0)", groupAdminId, groupRegistrationReturn.GroupId, userId)
	if err != nil {
		return err
	}
	return nil
}

func GroupInviteRegister(userId int, groupInviteMembershipReturn model.GroupInviteMembershipReturn) error {
	_, err := database.Db.Exec("INSERT INTO notif (target_user, id_one, id_two, notif_type) VALUES (?, ?, ?, 1)", groupInviteMembershipReturn.Target, groupInviteMembershipReturn.GroupId, userId)
	if err != nil {
		return err
	}
	return nil
}

func GroupMemberFetcher(claims jwt.MapClaims, groupId string) (bool, error) {
	var useless int

	err := database.Db.QueryRow(`
	SELECT DISTINCT users_group.user_id
	FROM users_group
	WHERE users_group.id = ?
	AND users_group.user_id = ?
	`, groupId, claims["id"]).Scan(&useless)

	if err == nil {
		return true, nil
	}

	err = database.Db.QueryRow(`
	SELECT DISTINCT group_members.user_id
	FROM group_members
	WHERE group_members.user_id = ?
	AND group_members.group_id = ?
	`, claims["id"], groupId).Scan(&useless)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func GroupMembersFetcher(userId int, groupId string) ([]model.GroupMembers, error) {
	var groupMembers model.GroupMembers
	var groupMembersList []model.GroupMembers

	rows, err := database.Db.Query(`
	SELECT DISTINCT group_members.*, users.nickname
	FROM group_members
	LEFT JOIN users ON users.id = group_members.user_id
	WHERE group_members.group_id = ?
	AND group_members.user_id != ?
	`, groupId, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&groupMembers.Id,
			&groupMembers.UserId,
			&groupMembers.GroupId,
			&groupMembers.UserName,
		)
		if err != nil {
			return nil, err
		}
		groupMembersList = append(groupMembersList, groupMembers)
	}
	return groupMembersList, nil
}

func ConfirmGroupMembership(askingUser int, groupId int) error {
	_, err := database.Db.Exec(`INSERT INTO group_members (user_id, group_id) VALUES(?,?)`, askingUser, groupId)

	return err
}
