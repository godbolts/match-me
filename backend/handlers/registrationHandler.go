package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/auth"
	"match_me_backend/db"
	"net/http"
	"regexp"
)

type RegLoginUser struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

type ErrorResponse struct {
	Error string `json:"error"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var user RegLoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Invalid register payload: %v", err)
		return
	}
	if user.Email == "" {
		sendErrorResponse(w, "Email is empty", http.StatusBadRequest)
		log.Printf("Email is empty: %v", err)
		return
	}
	if user.Password == "" {
		sendErrorResponse(w, "Password is empty", http.StatusBadRequest)
		log.Printf("Password is empty: %v", err)
		return
	}
	if len(user.Password) < 6 {
		sendErrorResponse(w, "Password should be at least 6 characters", http.StatusBadRequest)
		log.Printf("Password should be at least 6 characters: %v", err)
		return
	}
	if !emailRegex.MatchString(user.Email) {
		sendErrorResponse(w, "Not a valid email", http.StatusBadRequest)
		log.Printf("Not a valid email: %v", err)
		return
	}
	userExists, err := db.GetUserByEmail(user.Email)
	if err != nil && err.Error() != "no user found with that email" {
		sendErrorResponse(w, "Error checking user existence", http.StatusInternalServerError)
		log.Printf("no user found with that email: %v", err)
		return
	}
	if userExists != nil {
		sendErrorResponse(w, "User with this email already exists", http.StatusConflict)
		log.Printf("User with this email already exists: %v", err)
		return
	}
	hashedPassword, err := auth.HashPassword(user.Password)
	if err != nil {
		sendErrorResponse(w, "Failed to hash password", http.StatusInternalServerError)
		log.Printf("Failed to hash password: %v", err)
		return
	}
	err = db.SaveUser(user.Email, hashedPassword)	
	if err != nil {
		sendErrorResponse(w, "Error saving user", http.StatusInternalServerError)
		log.Printf("Error saving user: %v", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "User successfully registered"})
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: message})
}
