package models

import "time"

type LightProfileInformation struct {
	Username string `json:"username"`
	Picture  string `json:"profile_picture"`
}
type ProfileInformation struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Created   time.Time `json:"created_at"`
	IsOnline  bool      `json:"is_online"`
	City      string    `json:"user_city"`
	Nation    string    `json:"user_nation"`
	Region    string    `json:"user_region"`
	About     string    `json:"about_me"`
	Birthdate time.Time `json:"birthdate"`
	Age       string    `json:"age"`
	Picture   string    `json:"profile_picture"`
}

type GetUsername struct {
	Username string `json:"username"`
}

type GetCity struct {
	City   string `json:"user_city"`
	Nation string `json:"user_nation"`
	Region string `json:"user_region"`
}

type GetAbout struct {
	About string `json:"about_me"`
}

type GetBirthdate struct {
	Birthdate time.Time `json:"birthdate"`
	Age       string    `json:"username"`
}
