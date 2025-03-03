package models

import "time"

type User struct {
	ID               string    `json:"id"`
	Uuid             string    `json:"uuid"`
	Email            string    `json:"email"`
	PasswordHash     string    `json:"password_hash"`
	CreatedAt        time.Time `json:"created_at"`
	UserCity         string    `json:"user_city"`
	UserNation       string    `json:"user_nation"`
	UserRegion       string    `json:"user_region"`
	Latitude         float64   `json:"latitude"`
	Longitude        float64   `json:"longitude"`
	RegisterLocation string    `json:"register_location"`
	BrowsLocation    string    `json:"brows_location"`
	IsOnline 		 bool      `json:"is_online"`
}
