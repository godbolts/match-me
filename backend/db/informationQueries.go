package db

import (
	"fmt"
	"log"
	"time"
)

func SetUsername(userID, username string) error {
	userQuery := "UPDATE profiles SET username = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, username, userID)
	if err != nil {
		log.Printf("Error updating username for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update username: %w", err)
	}
	return nil
}

func GetUsername(userID string) string {
	var queriedUsername string
	userQuery := "SELECT username FROM profiles WHERE uuid = $1"
	err := DB.QueryRow(userQuery, userID).Scan(queriedUsername)
	if err != nil {
		log.Printf("Error updating username for uuid=%s: %v", userID, err)
		return ""
	}
	return queriedUsername
}

func SetCity(userID, nation, region, city string, latitude, longitude float64) error {
	log.Println("Setting city", city, "for user", userID, "with latitude", latitude, "and longitude", longitude)
	userQuery := `
	UPDATE users 
	SET 
		user_city = $1, 
		user_nation = $2,
		user_region = $3,
		latitude = $4,
		longitude = $5
	WHERE uuid = $6
	`
	_, err := DB.Exec(userQuery, city, nation, region, latitude, longitude, userID)
	if err != nil {
		log.Printf("Error updating city for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update city: %w", err)
	}
	return nil
}

func GetCity(userID string) string { //Need to work on coordinates
	var queriedCity string
	userQuery := "SELECT city FROM users WHERE uuid = $1"
	err := DB.QueryRow(userQuery, userID).Scan(queriedCity)
	if err != nil {
		log.Printf("Error updating city for uuid=%s: %v", userID, err)
		return ""
	}
	return queriedCity
}

func SetAbout(userID, about string) error {
	userQuery := "UPDATE profiles SET about_me = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, about, userID)
	if err != nil {
		log.Printf("Error updating about me for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update the about me: %w", err)
	}
	return nil
}

func GetAbout(userID string) string {
	var queriedAbout string
	userQuery := "SELECT about_me FROM profiles WHERE uuid = $1"
	err := DB.QueryRow(userQuery, userID).Scan(queriedAbout)
	if err != nil {
		log.Printf("Error updating city for uuid=%s: %v", userID, err)
		return ""
	}
	return queriedAbout
}

func SetBirthdate(userID string, birthdate time.Time) error {
	userQuery := "UPDATE profiles SET birthdate = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, birthdate, userID)
	if err != nil {
		log.Printf("Error updating birthdate for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update birthdate: %w", err)
	}
	return nil
}

func GetBirthdate(userID string) (string, int) {
	var queriedBirthdate string
	userQuery := "SELECT birthdate FROM profiles WHERE uuid = $1"
	err := DB.QueryRow(userQuery, userID).Scan(queriedBirthdate)
	if err != nil {
		log.Printf("Error updating city for uuid=%s: %v", userID, err)
		return "", 0
	}
	var age int
	if queriedBirthdate != "" {
		birthdayTime, err := time.Parse("2006-01-02T15:04:05Z", queriedBirthdate)
		if err != nil {
			log.Printf("Error parsing birthday for user_id %s: %v", userID, err)
			return "", 0
		}
		currentYear := time.Now().Year()
		age = currentYear - birthdayTime.Year()
		if birthdayTime.After(time.Now().AddDate(-age, 0, 0)) {
			age--
		}
	} else {
		age = 0
	}
	return queriedBirthdate, age
}
