package models

import "time"

// keeps track of all chats

type Chats struct {
	ID        int       `json:"id"`
	UserID1   string       `json:"userId1"`
	UserID2   string       `json:"userId2"`
	CreatedAt time.Time `json:"createdAt"`
}
