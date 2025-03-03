package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"net/http"
	"strings"
)

// LocationPayload defines the structure for the incoming location data
type LocationPayload struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func BrowserHandler(w http.ResponseWriter, r *http.Request) {
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
	var payload LocationPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("Error decoding request body: %v", err)
		return
	}
	if payload.Latitude == 0 || payload.Longitude == 0 {
		http.Error(w, "Latitude and Longitude cannot be zero", http.StatusBadRequest)
		log.Printf("Error: Latitude or Longitude is zero")
		return
	}
	Longitude := fmt.Sprintf("%f", payload.Longitude)
	Latitude := fmt.Sprintf("%f", payload.Latitude)
	err = db.SetBrowser(currentUserID, Longitude, Latitude)
	if err != nil {
		http.Error(w, "Error setting the Browser Location", http.StatusInternalServerError)
		log.Printf("Error setting the Browser Location: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Browser location successfully stored!"})
}
