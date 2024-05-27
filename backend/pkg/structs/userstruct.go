package socialnetwork

type UserInfo struct {
	Id             int    `json:"id"`
	Nickname       string `json:"nickname"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DateOfBirth    string `json:"date_of_birth"`
	About_me       string `json:"about_me"`
	Avatar         string `json:"avatar"`
	Public_Private bool   `json:"public_private"`
	CreatedAt      string `json:"created_at"`
	FollowStatus   int    `json:"follow_status"`
}

type AvatarAndUsername struct {
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
