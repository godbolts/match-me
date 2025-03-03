package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"

	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"match_me_backend/models"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type UserResponse struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	ProfilePicture string `json:"profile_picture"`
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	// Check if the user is authorized
	requesterID, err := GetCurrentUserID(r)
	if err != nil {
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("User/Profile not found for uuid=%s: %v", requesterID, err)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	user, err := db.GetUserByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "User/Profile not found", http.StatusNotFound)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("User/Profile not found for uuid=%s: %v", userID, err)
		} else {
			// http.Error(w, "Internal server error", http.StatusInternalServerError)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("Error fetching user/profile for uuid=%s: %v", userID, err)
		}
		return
	}
	user.Email = ""
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUserProfileHandler(w http.ResponseWriter, r *http.Request) {

	requesterID, err := GetCurrentUserID(r)
	if err != nil {
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("User/Profile not found for uuid=%s: %v", requesterID, err)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	profile, err := db.GetUserProfileByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "User/Profile not found", http.StatusNotFound)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("User/Profile not found for uuid=%s: %v", userID, err)
		} else {
			// http.Error(w, "Internal server error", http.StatusInternalServerError)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("Error fetching user/profile for uuid=%s: %v", userID, err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func GetUserBioHandler(w http.ResponseWriter, r *http.Request) {

	requesterID, err := GetCurrentUserID(r)
	if err != nil {
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("User/Profile not found for uuid=%s: %v", requesterID, err)
		return
	}

	vars := mux.Vars(r)
	userID := vars["id"]
	profile, err := db.GetUserBioByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			// http.Error(w, "User/Profile not found", http.StatusNotFound)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("User/Profile not found for uuid=%s: %v", userID, err)
		} else {
			// http.Error(w, "Internal server error", http.StatusInternalServerError)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("Error fetching user/profile for uuid=%s: %v", userID, err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func GetMeBioHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		// http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	profile, err := db.GetUserBioByID(currentUserID)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User/Profile not found", http.StatusNotFound)
			log.Printf("User/Profile not found for uuid=%s: %v", currentUserID, err)
		} else {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			log.Printf("Error fetching user/profile for uuid=%s: %v", currentUserID, err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)

}

func GetCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	var user *models.ProfileInformation
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		// http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	user, err = db.GetUserInformation(currentUserID)
	if err != nil {
		if err.Error() == "user not found" {
			// http.Error(w, "User not found", http.StatusNotFound)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("User with ID %v not found", currentUserID)
		} else {
			// http.Error(w, "Internal server error", http.StatusInternalServerError)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("Error fetching user information: %v", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func AuthorizationHandler(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	_, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		// http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Error validating token: %v", err)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("Authorization successful")
}

func GetLightCurrentUserHandler(w http.ResponseWriter, r *http.Request) {
	var user *models.LightProfileInformation
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	user, err = db.GetLightUserInformation(currentUserID)
	if err != nil {
		if err.Error() == "user not found" {
			// http.Error(w, "User not found", http.StatusNotFound)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("User with ID %v not found", currentUserID)
		} else {
			// http.Error(w, "Internal server error", http.StatusInternalServerError)
			http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
			log.Printf("Error fetching user information: %v", err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetCurrentUserID(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		log.Printf("Unauthorized: Missing or invalid token")
		return "", errors.New("unauthorized: Missing or invalid token in header")

	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		return "", err
	}
	return currentUserID, nil
}

func GetCurrentUserUUID(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		// http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Unauthorized: Missing or invalid token")
		return
	}
	token := strings.TrimPrefix(authHeader, "Bearer ")
	currentUserID, err := auth.ExtractUserIDFromToken(token)
	if err != nil {
		// http.Error(w, "Unauthorized: Invalid or expired token", http.StatusUnauthorized)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		log.Printf("Error extracting user ID from token: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(currentUserID)
}
