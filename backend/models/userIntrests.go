package models

// This data structure stores the user and what interest they have.
// The more of the interest  match the higher the score.

// From the front end we send back the user ID and selected InterestID.

type UserInterests struct {
	ID         int    `json:"id"`
	UserID     string `json:"userId"`
	InterestID int    `json:"interestId"`
	Status     string `json:"status"`
}
