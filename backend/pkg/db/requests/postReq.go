package socialnetwork

import (
	"database/sql"

	database "socialnetwork/pkg/db/sqlite"
	model "socialnetwork/pkg/structs"
)

func PostInserter(post model.PostReturn, groupId string, client int) error {
	
	var err error

	if len(groupId) == 0 {
		_, err = database.Db.Exec("INSERT INTO posts (title, body, img, author_id, view_status) VALUES (?, ?, ?, ?, ?)", post.Title, post.Body, post.ImgName, client, post.ViewStatus)
		if err != nil {
			return err
		}

		var postId int
		//get last inserted
		err = database.Db.QueryRow("SELECT last_insert_rowid() FROM posts").Scan(&postId)
		if err != nil {
			return err
		}

		for _, v := range post.Followers {
			if v.Checked {
				_, err = database.Db.Exec("INSERT INTO view_post (user_id, post_id) VALUES (?, ?)", v.Id, postId)
				if err != nil {
					return err
				}
			}
		}
	} else {
		_, err = database.Db.Exec("INSERT INTO posts (title, body, img, author_id, group_id, view_status) VALUES (?, ?, ?, ?, ?, ?)", post.Title, post.Body, post.ImgName , client, groupId, 0)
		if err != nil {
			return err
		}
	}
	return nil
}

func PostFetcher(groupId string, client int) ([]model.Post, error){

	var rows *sql.Rows
	var post model.Post
	var err error
	var postList []model.Post

	if len(groupId) == 0 {
		rows, err = database.Db.Query(`
		SELECT DISTINCT posts.*, users.nickname
		FROM posts
		LEFT JOIN view_post ON view_post.post_id = posts.id
		LEFT JOIN users ON users.id = posts.author_id
		LEFT JOIN follows ON follows.follower_user_id = ? AND follows.following_user_id = posts.author_id AND follows.follow_status = 1
		WHERE posts.group_id IS NULL
		AND (
			posts.author_id = ?
			OR view_status = 0
			OR (
				view_status = 2 AND (
					follows.follower_user_id = ?
				)
			)
			OR (
				view_status = 1 AND (
					view_post.user_id = ?
				)
			)
		)
		ORDER BY posts.id DESC
		`, client, client, client, client)
		if err != nil {
			return nil, err
		}
	} else if len(groupId) > 5  {
		if groupId[:6] == "profil" {
			userSelectedId, err := GetUserIdFromNick(groupId[6:])
			if err != nil {
				return nil, err
			}
			rows, err = database.Db.Query(`
			SELECT DISTINCT posts.*, users.nickname
			FROM posts
			LEFT JOIN view_post ON view_post.post_id = posts.id
			LEFT JOIN users ON users.id = posts.author_id
			LEFT JOIN follows ON follows.follower_user_id = ? AND follows.following_user_id = posts.author_id AND follows.follow_status = 1
			WHERE posts.author_id = ?
			AND (
				posts.author_id = ?
				OR view_status = 0
				OR (
					view_status = 2 AND (
						follows.follower_user_id = ?
					)
				)
				OR (
					view_status = 1 AND (
						view_post.user_id = ?
					)
				)
			)
			ORDER BY posts.id DESC
			`, client, userSelectedId, client, client, client)
			if err != nil {
				return nil, err
			}
		}
	} else {
		rows, err = database.Db.Query(`
		SELECT DISTINCT posts.*, users.nickname
		FROM posts
		LEFT JOIN users ON  users.id = posts.author_id
		WHERE posts.group_id = ?
		ORDER BY posts.id DESC
		`, groupId)
		if err != nil {
			return nil, err
		}
	}

	for rows.Next() {
		err = rows.Scan(
			&post.Id,
			&post.Title,
			&post.Body,
			&post.Img,
			&post.AuthorId,
			&post.GroupId,
			&post.ViewStatus,
			&post.CreatedAt,
			&post.AuthorName,
		)
		if err != nil {
			return nil, err
		}
		postList = append(postList, post)
	}
	return postList, nil
}
