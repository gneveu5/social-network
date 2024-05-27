package socialnetwork

type Notification struct {
	Id        int `json:"id"`
	IdOne     int `json:"idOne"`
	IdTwo     int `json:"idTwo"`
	NotifType int `json:"notifType"`
}

type GroupMembershipRequestNotif struct {
	Notification `json:"notification"`
	GroupName    string `json:"groupName"`
	UserName     string `json:"userName"`
}

type GroupInviteNotif struct {
	Notification `json:"notification"`
	GroupName    string `json:"groupName"`
	UserName     string `json:"userName"`
}

type FollowRequestNotif struct {
	Notification `json:"notification"`
	UserName     string `json:"userName"`
}

type EventNotif struct {
	Notification `json:"notification"`
	GroupName    string `json:"groupName"`
	EventName    string `json:"eventName"`
}

type Notifications struct {
	GroupInvite            []GroupInviteNotif            `json:"groupInvite"`
	GroupMembershipRequest []GroupMembershipRequestNotif `json:"groupMembershipRequest"`
	FollowRequest          []FollowRequestNotif          `json:"followRequest"`
	Event                  []EventNotif                  `json:"event"`
}
