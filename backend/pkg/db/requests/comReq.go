package socialnetwork

import (

	database "socialnetwork/pkg/db/sqlite"
	model "socialnetwork/pkg/structs"
)

func CommentInserter(comment model.CommentPosting, client int) error {

	_, err := database.Db.Exec("INSERT INTO post_reply (user_id, post_id, body, img) VALUES (?, ?, ?, ?)", client, comment.PostId, comment.Message, comment.ImgName)

	if err != nil {
		return err
	}

	return nil
}
	
func CommentFetcher(postId string) ([]model.Comment, error) {

	var comment model.Comment
	var commentList []model.Comment

	rows, err := database.Db.Query(`
	SELECT post_reply.*, users.nickname
	FROM post_reply
	LEFT JOIN users ON  users.id = post_reply.user_id
	WHERE post_id = ?
	ORDER BY id DESC
	`, postId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(
			&comment.Id,
			&comment.UserId,
			&comment.PostId,
			&comment.Body,
			&comment.Img,
			&comment.CreatedAt,
			&comment.UserName,
		)
		if err != nil {
			return nil, err
		}
		commentList = append(commentList, comment)
	}
	return commentList, nil
}
