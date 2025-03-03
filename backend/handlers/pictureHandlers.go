package handlers

import (
	"fmt"
	"io"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func PostProfilePictureHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	file, _, err := r.FormFile("profilePic") // "profilePic" should be the name of the form field
	if err != nil {
		currentPic := db.GetPicturePath(currentUserID)
		if currentPic != "" {
			log.Printf("Error setting picture but one already exists so nothing will be changed.")
			return
		} else {
			err = db.SetPicturePath(currentUserID, "default_profile_pic.png")
			if err != nil {
				http.Error(w, "Error setting the default profile picture path", http.StatusInternalServerError)
				log.Printf("Error setting the default profile picture path: %v", err)
				return
			}
		}
		return
	}
	defer file.Close()
	randomFileName := currentUserID + ".jpeg"
	dir := "../frontend/src/components/Assets/ProfilePictures" // Change this to the desired directory
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			http.Error(w, "Unable to create directory", http.StatusInternalServerError)
			log.Printf("Error creating directory: %v", err)
			return
		}
	}
	filePath := filepath.Join(dir, randomFileName)
	destFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to save file", http.StatusInternalServerError)
		log.Printf("Error creating file: %v", err)
		return
	}
	defer destFile.Close()
	_, err = io.Copy(destFile, file)
	if err != nil {
		http.Error(w, "Unable to save file content", http.StatusInternalServerError)
		log.Printf("Error copying file content: %v", err)
		return
	}
	err = db.SetPicturePath(currentUserID, randomFileName)
	if err != nil {
		http.Error(w, "Error setting the username", http.StatusInternalServerError)
		log.Printf("Error setting the username: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Profile picture uploaded successfully! Saved as %s", randomFileName)))
}

func PostProfileRPictureRemoveHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	dir := "../frontend/src/components/Assets/ProfilePictures" 
	profilePicFileName := currentUserID + ".jpeg"             
	filePath := filepath.Join(dir, profilePicFileName)
	if _, err := os.Stat(filePath); err == nil {
		err := os.Remove(filePath)
		if err != nil {
			http.Error(w, "Unable to delete profile picture", http.StatusInternalServerError)
			log.Printf("Error deleting file: %v", err)
			return
		}
		log.Printf("Deleted profile picture: %s", filePath)
	} else if !os.IsNotExist(err) {
		http.Error(w, "Unable to check file existence", http.StatusInternalServerError)
		log.Printf("Error checking file existence: %v", err)
		return
	}
	err = db.SetPicturePath(currentUserID, "default_profile_pic.png")
	if err != nil {
		http.Error(w, "Unable to update profile picture path in database", http.StatusInternalServerError)
		log.Printf("Error updating database: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Profile picture removed and reset to default successfully"))
	log.Printf("Profile picture removed and reset for user ID: %s", currentUserID)
}