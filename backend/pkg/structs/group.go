package socialnetwork

import (
	"time"
)

type Group struct {
	Id			int
	Title		string
	Description	string
	UserId		int
	CreatedAt	time.Time
	AdminName	string
}

type GroupAskMembershipReturn struct {
	GroupId	string
}

type GroupInviteMembershipReturn struct {
	Target	int
	GroupId int
}

type GroupCreationReturn struct {
	Title		string
	Description string
}

type GroupMembers struct {
	Id			int
	UserId		int
	GroupId		int
	UserName	string
}
