package socialnetwork

type Follow struct {
	ID       int    `json:"id"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}
