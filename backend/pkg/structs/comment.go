package socialnetwork

import (
	"database/sql"
	"time"
	"mime/multipart"
)

type Comment struct {
	Id			int
	Body		string
	Img			sql.NullString
	UserId		int
	PostId		int
	CreatedAt	time.Time
	UserName	string
}

type CommentPosting struct {
	PostId	string
	Message string
	ImgFile multipart.File
	ImgName	string
}

type CommentFetchReturn struct {
	PostId	int
}
