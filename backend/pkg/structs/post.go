package socialnetwork

import (
	"database/sql"
	"mime/multipart"
	"time"
)

type Post struct {
	Id			int
	Title		string
	Body		string
	Img			sql.NullString
	AuthorId	sql.NullInt64
	GroupId		sql.NullInt64
	ViewStatus	int
	CreatedAt	time.Time
	AuthorName	string
}

type PostReturn struct {
	Title 		string
	Body  		string
	Followers	[]FollowerPost
	ImgFile   	multipart.File
	ImgName		string
	GroupId 	string
	ViewStatus  string
}

type FollowerPost struct {
    Avatar		string
	Checked		bool
	Id			int
	Nickname	string
}
