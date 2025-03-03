package models

import "time"

// The friend status is gotten from this structure, also blocks or deletes.
// We get the friends list by

type UserConnections struct {
	ID        int       `json:"id"`
	UserID1   string    `json:"userId1"`
	UserID2   string    `json:"userId2"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
}

type ConnectionRequest struct {
	UserID string `json:"user_id"`
}
