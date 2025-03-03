package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUserInterests(w http.ResponseWriter, r *http.Request) {
	userID, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error:", err)
	}
	userInterestIDs, err := db.GetAllUserInterestIDs(userID)
	if err != nil {
		log.Println("Error in GetUserInterestsId's", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userInterestIDs)
}

func GetIDUserInterests(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	userInterestIDs, err := db.GetAllUserInterestIDs(userID)
	if err != nil {
		log.Println("Error in GetUserInterestsId's", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userInterestIDs)
}

func GetInterests(w http.ResponseWriter, r *http.Request) {
	interests, err := db.GetInterestResponseBody()
	if err != nil {
		log.Println("Error getting interest response ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(interests)
}

func UpdateUserInterest(w http.ResponseWriter, r *http.Request) {
	type InterestRequest struct {
		InterestID int `json:"interestId"`
	}
	userID, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id:", err)
	}
	var request InterestRequest
	err = json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	interestID := request.InterestID
	err = db.AddInterestToUser(interestID, userID)
	if err != nil {
		log.Println("Error adding interest to user:", err)
		http.Error(w, "Error adding interest to user", http.StatusInternalServerError)
		return
	}
	db.UpdateMatchScoreForUser(userID)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Interest added successfully"})
}
