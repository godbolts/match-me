package handlers

import (
	"encoding/json"
	"log"
	"match_me_backend/db"
	"time"

	"net/http"
)

//to test any GET function use postman and run localhost:4000/test

func GetDemoUsers(w http.ResponseWriter, r *http.Request) {
	// cycles to set DB pause for every 1000 demo users spawned. 
	cycles  := 0
	if db.InitDemoUsers() {
		userMatches, err := db.GetAllUserMatches()
		if err != nil {
			log.Println("Error getting user matches:", err)
		}
		for _, userMatch := range userMatches {
			db.CalculateUserDistance(userMatch.UserID1, userMatch.UserID2)
			db.CalculateMatchScore(userMatch.UserID1, userMatch.UserID2)
			if cycles == 1000 {
				time.Sleep(1 * time.Second)
				log.Println(" demo users spawned")
				cycles = 0
			}
			cycles++
		}
	}
	successMessage := "Demo bots spawned and are on the loose!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successMessage)
}
