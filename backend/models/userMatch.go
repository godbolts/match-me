package models

import (
	"time"
)
type UsersMatches struct {
	ID         int       `json:"id"`
	UserID1    string    `json:"user_id_1"`
	UserID2    string    `json:"user_id_2"`
	MatchScore int       `json:"match_score"`
	Status     string    `json:"status"`
	Requester  string    `json:"requester"`  //  new, shown, connected, blocked, deleted
	ModifiedAt time.Time `json:"modifiedAt"` // last time the match was updated
	CreatedAt  time.Time `json:"createdAt"`  // first time the match was created
	Distance   float64   `json:"distance"`   // distance between the two users
}
