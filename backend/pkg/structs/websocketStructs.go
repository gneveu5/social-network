package socialnetwork

import (
	"github.com/gorilla/websocket"
)

type ConnectedUserData struct {
	Conns  []*websocket.Conn // one user can be logged on different devices, store a websocket for each one
	UserId int
	Name   string
}

type WebsocketMessage struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

// single contact data
type Contact struct {
	ContactId int                  `json:"contactId"`
	Name      string               `json:"name"`
	History   []ChatHistoryElement `json:"history"`
}

type PrivateContact struct {
	Contact
	Connected bool `json:"connected"`
}

// used to send contact list to user when he connects. type: "init-contact"
type ContactListMsg struct {
	Private []PrivateContact `json:"private"`
	Group   []Contact        `json:"group"`
}

type RemoveContact struct {
	Id int `json:"id"`
}

// holds single message data. types : "pvt-msg" or "pvt-confirm"
type ChatMessage struct {
	To  int    `json:"to"` // recipient id in database
	Txt string `json:"txt"`
}

type ChatMessageToClient struct {
	ChatMessage
	Id   int    `json:"id"`
	From string `json:"from"`
	Date int64  `json:"date"`
}

type ChatHistoryElement struct {
	Id   int    `json:"id"`
	From string `json:"from"`
	Date int64  `json:"date"`
	Txt  string `json:"txt"`
}

type HistoryRequest struct {
	Id   int `json:"id"`   // group_id or user_id, if group or private conversation
	Last int `json:"last"` // last message id
}

type HistoryResponse struct {
	To   int                  `json:"to"`
	Hist []ChatHistoryElement `json:"hist"`
}

// used to warn connected users that someone else connected/disconnected. types: "connect" or "disconnect"
type UpdateConnectedStatus struct {
	UserId int `json:"userId"`
}

type NotificationCount struct {
	Count int `json:"count"`
}
