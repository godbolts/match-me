package models

import "time"

//message history

type Messages struct {
	ID        int       `json:"id"`
	ChatID    int       `json:"chatId"`
	SenderID  int       `json:"senderId"`
	Messages  string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
