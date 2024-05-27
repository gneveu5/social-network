package socialnetwork

import "mime/multipart"

type User struct {
	ID            uint
	NickName      string
	Email         string
	Password      string
	FirstName     string
	LastName      string
	DateOfBirth   string // le format est peut etre à changer
	AboutMe       string
	PublicPrivate bool
	Avatar        string
	CreatedAt     string
	HasValidated  bool
}

type SingUpRequest struct {
	Email       string
	FirstName   string
	LastName    string
	DateOfBirth string // le format est peut etre à changer
	NickName    string
	AboutMe     string
	AvatarFile  multipart.File
	AvatarName  string
	Password    string
}
type LoginResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
