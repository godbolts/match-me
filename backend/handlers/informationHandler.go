package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type PostUsernameRequest struct {
	Username string `json:"username"`
}

type PostAboutRequest struct {
	About string `json:"about"`
}

type PostBirthdateRequest struct {
	Birthdate time.Time `json:"birthdate"`
}

func PostUsername(w http.ResponseWriter, r *http.Request) {
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
	var body PostUsernameRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.Username == "" {
		http.Error(w, "Username cannot be empty", http.StatusBadRequest)
		log.Printf("Error: Username cannot be empty")
		return
	}
	err = db.SetUsername(currentUserID, body.Username)
	if err != nil {
		http.Error(w, "Error setting the username", http.StatusInternalServerError)
		log.Printf("Error setting the username: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Username successfully registered"})
}

type PostCityNameRequest struct {
	City   string `json:"city"`
	Nation string `json:"country"`
	Region string `json:"state"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
}

func PostCity(w http.ResponseWriter, r *http.Request) {
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
	var body PostCityNameRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.City == "" || body.Region == "" || body.Nation == "" {
		http.Error(w, "City details cannot be empty", http.StatusBadRequest)
		log.Printf("Error: City, Longitude, or Latitude cannot be empty")
		return
	}
	latitude64, err := strconv.ParseFloat(body.Latitude, 64)
	longitude64, err := strconv.ParseFloat(body.Longitude, 64)
	err = db.SetCity(currentUserID, body.Nation, body.Region, body.City, latitude64, longitude64)
	if err != nil {
		http.Error(w, "Error setting the City", http.StatusInternalServerError)
		log.Printf("Error setting the City: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "City successfully registered"})
}

func PostAbout(w http.ResponseWriter, r *http.Request) {
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
	var body PostAboutRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.About == "" {
		http.Error(w, "About cannot be empty", http.StatusBadRequest)
		log.Printf("Error: About cannot be empty for userID=%s", currentUserID)
		return
	}
	err = db.SetAbout(currentUserID, body.About)
	if err != nil {
		http.Error(w, "Error setting the About", http.StatusInternalServerError)
		log.Printf("Error setting the About for userID=%s: %v", currentUserID, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "About successfully registered"})
}

func PostBirthdate(w http.ResponseWriter, r *http.Request) {
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
	var body PostBirthdateRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.Birthdate.IsZero() {
		http.Error(w, "Birthdate cannot be empty", http.StatusBadRequest)
		log.Printf("Error: Birthdate cannot be empty for userID=%s", currentUserID)
		return
	}
	err = db.SetBirthdate(currentUserID, body.Birthdate)
	if err != nil {
		http.Error(w, "Error setting the birthdate", http.StatusInternalServerError)
		log.Printf("Error setting the birthdate for userID=%s: %v", currentUserID, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Birthdate successfully registered"})
}
