package models

import (
	"time"
)

//TODO: add correct table structure.

type Profiles struct {
	ID             int       `json:"id"`
	UUID           string    `json:"uuid"`
	Username       string    `json:"username"`
	AboutMe        string    `json:"about_me"`
	ProfilePicture string    `json:"profile_picture"`
	CreatedAt      time.Time `json:"created_at"`
	Birthdate      time.Time `json:"birthdate"`
}

// CREATE TABLE profiles (
//     id SERIAL PRIMARY KEY,
//     user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,  //If there are no foreign key constraints it will not cascade
//     username VARCHAR(20) NOT NULL UNIQUE,
//     about_me TEXT,
//     profile_picture TEXT,
//     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// );
