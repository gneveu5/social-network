package socialnetwork

import (
	"time"
)

type Event struct {
	Id			int
	Title		string
	Description	string
	EventDate	string
	UserId		int
	GroupId		int
	CreatedAt	time.Time
	CreatorName	string
}

type EventReturn struct {
	Title 		string
	Description string
	EventDate	string
	GroupId 	string
}

type EventRegistrationReturn struct {
	EventId	int
	Status	string
}

type EventAttendees struct {
	Id			int
	UserId		int
	EventId		int
	AttendeeName string
}
