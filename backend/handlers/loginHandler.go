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

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var user RegLoginUser
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		sendErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		log.Printf("Invalid login request: %v", err)
		return
	}
	if user.Password == "" || (user.Email == "" && user.Username == "") {
		sendErrorResponse(w, "Password and either Email or Username are required", http.StatusBadRequest)
		log.Println("Password and either Email or Username are required")
		return
	}
	var existingUser *models.User
	if user.Username != "" {
		existingUser, err = db.GetUserByUsername(user.Username)
	} else {
		existingUser, err = db.GetUserByEmail(user.Email)
	}
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
			log.Println("Email or password is incorrect")
			return
		}
		sendErrorResponse(w, "Error checking user existence", http.StatusInternalServerError)
		log.Printf("Error checking user existence: %v", err)
		return
	}
	if existingUser == nil || !auth.ComparePasswords(existingUser.PasswordHash, user.Password) {
		sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
		log.Println("Email or password is incorrect")
		return
	}
	token, err := auth.GenerateJWT(existingUser.ID)
	if err != nil {
		sendErrorResponse(w, "Error generating JWT", http.StatusInternalServerError)
		log.Printf("Error generating JWT: %v", err)
		return
	}
	if token == "" {
		sendErrorResponse(w, "No token", http.StatusInternalServerError)
		log.Println("Error token is empty")
		return
	}
	log.Println("Starting match Updated", )
	err = db.UpdateMatchScoreForUser(existingUser.ID)
	if err != nil {
		log.Println("Error updating all user scores", err)
	}
	log.Println("Starting to set User online", )
	err = db.SetUserOnlineStatus(existingUser.ID, true)
	if err != nil {
		log.Println("Error setting user online status", err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}


func LoginAPIHandler(w http.ResponseWriter, r *http.Request) {
    email := r.URL.Query().Get("email")
    password := r.URL.Query().Get("password")
    if password == "" || email == "" {
        sendErrorResponse(w, "Email and password are required", http.StatusBadRequest)
        log.Println("Email and password are required")
        return
    }
    existingUser, err := db.GetUserByEmail(email)
    if err != nil {
        if errors.Is(err, db.ErrUserNotFound) {
            sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
            log.Println("Email or password is incorrect")
            return
        }
        sendErrorResponse(w, "Error checking user existence", http.StatusInternalServerError)
        log.Printf("Error checking user existence: %v", err)
        return
    }
    if existingUser == nil || !auth.ComparePasswords(existingUser.PasswordHash, password) {
        sendErrorResponse(w, "Email or password is incorrect", http.StatusUnauthorized)
        log.Println("Email or password is incorrect")
        return
    }
    token, err := auth.GenerateJWT(existingUser.ID)
    if err != nil {
        sendErrorResponse(w, "Error generating JWT", http.StatusInternalServerError)
        log.Printf("Error generating JWT: %v", err)
        return
    }
    if token == "" {
        sendErrorResponse(w, "No token generated", http.StatusInternalServerError)
        log.Println("Error token is empty")
        return
    }
    err = db.UpdateMatchScoreForUser(existingUser.ID)
    if err != nil {
        log.Println("Error updating match score:", err)
    }
    err = db.SetUserOnlineStatus(existingUser.ID, true)
    if err != nil {
        log.Println("Error setting user online status:", err)
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := GetCurrentUserID(r)
	log.Println("userId after we got it from GetCurrentUserID(r)", userId)
	if err != nil {
		// sendErrorResponse(w, "Error getting user Id from token", http.StatusUnauthorized)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(false)
		log.Println("Error getting user Id from token:", err)
		return
	}
	err = db.SetUserOnlineStatus(userId, false)
	if err != nil {
		// sendErrorResponse(w, "Error setting user offline", http.StatusInternalServerError)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(false)
		log.Println("Error setting user offline:", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(true)
}

func GetOnlineStatus(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Unauthorized: Missing or invalid token", http.StatusUnauthorized)
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

	isOnline, err := db.GetUserOnlineStatus(currentUserID)
	if err != nil {
		log.Printf("Error fetching online status for user %s: %v", currentUserID, err)

		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]bool{"is_online": isOnline})
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func GetOtherOnlineStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	if userID == "" {
		http.Error(w, "Bad Request: Missing user ID", http.StatusBadRequest)
		log.Printf("Bad Request: Missing user ID")
		return
	}

	isOnline, err := db.GetUserOnlineStatus(userID)
	if err != nil {
		log.Printf("Error fetching online status for user %s: %v", userID, err)
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(isOnline)
	if err != nil {
		log.Printf("Error encoding response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
