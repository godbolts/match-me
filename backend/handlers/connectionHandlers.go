package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"match_me_backend/models"
	"net/http"
	"strings"
)

func PostFirstConnection(w http.ResponseWriter, r *http.Request) {
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
	var body models.ConnectionRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.UserID == "" {
		http.Error(w, "userID2 cannot be empty", http.StatusBadRequest)
		log.Printf("Error: userID2 cannot be empty")
		return
	}
	err = db.SetFirstConnection(currentUserID, body.UserID)
	if err != nil {
		http.Error(w, "Error setting the userID2", http.StatusInternalServerError)
		log.Printf("Error setting the userID2: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Connection successfully registered"})
}

func PostAcceptance(w http.ResponseWriter, r *http.Request) {
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
	var body models.ConnectionRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.UserID == "" {
		http.Error(w, "userID1 cannot be empty", http.StatusBadRequest)
		log.Printf("Error: userID1 cannot be empty")
		return
	}
	err = db.SetAccepted(currentUserID, body.UserID)
	if err != nil {
		http.Error(w, "Error setting the userID1", http.StatusInternalServerError)
		log.Printf("Error setting the userID1: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Connection successfully accepted"})
}

func PostBlock(w http.ResponseWriter, r *http.Request) {
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
	var body models.ConnectionRequest
	err = json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Error parsing request body: %v", err)
		return
	}
	if body.UserID == "" {
		http.Error(w, "userID1 cannot be empty", http.StatusBadRequest)
		log.Printf("Error: userID1 cannot be empty")
		return
	}
	err = db.Setblock(currentUserID, body.UserID)
	if err != nil {
		http.Error(w, "Error setting the userID1", http.StatusInternalServerError)
		log.Printf("Error setting the userID1: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Connection successfully blocked"})
}