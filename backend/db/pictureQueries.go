package db

import (
	"fmt"
	"log"
)

func SetPicturePath(userID, path string) error {
	userQuery := "UPDATE profiles SET profile_picture = $1 WHERE uuid = $2"
	_, err := DB.Exec(userQuery, path, userID)
	if err != nil {
		log.Printf("Error updating profile picture for uuid=%s: %v", userID, err)
		return fmt.Errorf("could not update the profile picture: %w", err)
	}
	return nil
}

func GetPicturePath(userID string) string {
	var queriedPicturePath string
	userQuery := "SELECT profile_picture FROM profiles WHERE uuid = $1"
	err := DB.QueryRow(userQuery, userID).Scan(&queriedPicturePath)
	if err != nil {
		log.Printf("Error returning picture path for uuid=%s: %v", userID, err)
		return ""
	}
	return queriedPicturePath
}
