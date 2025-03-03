package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	"match_me_backend/models"
	"net/http"
)

// Handle match requests from the front end

type MatchRequest struct {
	MatchId int `json:"match_id"`
}

func RemoveMatch(w http.ResponseWriter, r *http.Request) {
	var payload MatchRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	matchId := payload.MatchId
	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}
	_, err = GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	successMessage, err := db.UpdateUserMatchStatus(matchId, db.REMOVED)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

func RequestMatch(w http.ResponseWriter, r *http.Request) {
	var payload MatchRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	matchId := payload.MatchId
	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}
	requesterId, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	var successMessage string
	err = db.SetRequesterIdForMatch(requesterId, matchId)
	successMessage, err = db.UpdateUserMatchStatus(matchId, db.REQUESTED)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

func ConfirmMatch(w http.ResponseWriter, r *http.Request) {
	var payload MatchRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	matchId := payload.MatchId
	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}
	_, err = GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	successMessage, err := db.UpdateUserMatchStatus(matchId, db.CONNECTED)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

func BlockMatch(w http.ResponseWriter, r *http.Request) {
	var payload MatchRequest
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	matchId := payload.MatchId
	if matchId == 0 {
		http.Error(w, "Missing matchId", http.StatusBadRequest)
		return
	}
	_, err = GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	successMessage, err := db.UpdateUserMatchStatus(matchId, db.BLOCKED)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": successMessage,
	})
}

type MatchResponse struct {
	MatchID                int    `json:"match_id"`
	MatchScore             int    `json:"match_score"`
	Status                 string `json:"status"`          // Used to determine what to display in the front end
	MatchedUserID          string `json:"matched_user_id"` // UUID
	Requester              string `json:"requester"`
	MatchedUserName        string `json:"matched_user_name"`
	MatchedUserPicture     string `json:"matched_user_picture"`
	MatchedUserDescription string `json:"matched_user_description"`
	MatchedUserLocation    string `json:"matched_user_location"`
	IsOnline               bool   `json:"is_online"`
}

// Returns 10 best matches with the 'new' status
func GetMatches(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	userMatches, err := db.GetTenNewMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user matches:", err)
	}
	if err != nil {
		log.Println("Error getting user matches:", err)
	}
	var matches []MatchResponse
	var match MatchResponse
	var matchProfile *models.ProfileInformation
	var buddyID string
	for _, userMatch := range userMatches {
		if userMatch.UserID2 == userID1 {
			matchProfile, err = db.GetUserInformation(userMatch.UserID1)
			buddyID = userMatch.UserID1
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			matchProfile, err = db.GetUserInformation(userMatch.UserID2)
			buddyID = userMatch.UserID1
		}
		if err != nil {
			log.Println("Error getting user information:", err)
		}
		match.MatchID = userMatch.ID
		match.Status = userMatch.Status
		match.MatchScore = userMatch.MatchScore
		match.MatchedUserID = buddyID
		match.MatchedUserName = matchProfile.Username
		match.MatchedUserPicture = matchProfile.Picture
		match.MatchedUserDescription = matchProfile.About
		match.MatchedUserLocation = matchProfile.Nation
		match.IsOnline = matchProfile.IsOnline
		matches = append(matches, match)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matches)
}

func GetRequests(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	userMatches, err := db.GetRequestsMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user matches:", err)
	}
	if err != nil {
		log.Println("Error getting user matches:", err)
	}
	var matches []MatchResponse
	var match MatchResponse
	var matchProfile *models.ProfileInformation
	var buddyID string
	for _, userMatch := range userMatches {
		if userMatch.UserID2 == userID1 {
			matchProfile, err = db.GetUserInformation(userMatch.UserID1)
			buddyID = userMatch.UserID1
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			matchProfile, err = db.GetUserInformation(userMatch.UserID2)
			buddyID = userMatch.UserID1
		}
		if err != nil {
			log.Println("Error getting user information:", err)
		}
		if userID1 == userMatch.Requester {
			match.Requester = "true"
		} else {
			match.Requester = "false"
		}
		match.MatchID = userMatch.ID
		match.Status = userMatch.Status
		match.MatchScore = userMatch.MatchScore
		match.MatchedUserID = buddyID
		match.MatchedUserName = matchProfile.Username
		match.MatchedUserPicture = matchProfile.Picture
		match.MatchedUserDescription = matchProfile.About
		match.MatchedUserLocation = matchProfile.Nation
		match.IsOnline = matchProfile.IsOnline
		matches = append(matches, match)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(matches)
}

func GetRecommendationsHandler(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userMatches, err := db.GetAllUserMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user matches:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Println(`userMatches`, userMatches)
	var matchIDs []string
	var buddyID string
	for _, userMatch := range userMatches {
		if userMatch.UserID2 == userID1 {
			buddyID = userMatch.UserID1
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			buddyID = userMatch.UserID1
		}

		if err != nil {
			log.Println("Error getting user information:", err)
		}
		match := buddyID

		matchIDs = append(matchIDs, match)
	}

	// for _, userMatch := range userMatches {
	// 	matchProfile, err := db.GetUserInformation(userMatch.UserID2)
	// 	if err != nil {
	// 		log.Println("Error getting user information:", err)
	// 		continue
	// 	}
	// 	match := MatchResponse{
	// 		MatchedUserID:             userMatch.UserID1
	// 		M
	// MatchID:                userMatch.ID,
	// Status:                 userMatch.Status,
	// MatchScore:             userMatch.MatchScore,
	// MatchedUserName:        matchProfile.Username,
	// MatchedUserPicture:     matchProfile.Picture,
	// MatchedUserDescription: matchProfile.About,
	// MatchedUserLocation:    matchProfile.Nation,
	// IsOnline:               matchProfile.IsOnline,
	// 	}
	// 	matches = append(matches, match)
	// }
	// sort.Slice(matches, func(i, j int) bool {
	// 	return matches[i].MatchScore > matches[j].MatchScore
	// })
	// if len(matches) > 10 {
	// 	matches = matches[:10]
	// }
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(matchIDs)
}

func GetConnections(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		return
	}
	connectionsUUIDs, err := db.GetConnectionsID(userID1)
	if err != nil {
		log.Println("Error getting connections UUID from db:", err)
		http.Error(w, "HTTP Status: 404 (not found) ", http.StatusNotFound)
		return
		
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(connectionsUUIDs)
}

type BuddiesResponse struct {
	MatchID                int      `json:"match_id"`
	MatchScore             int      `json:"match_score"`
	Status                 string   `json:"status"`
	MatchedUserName        string   `json:"matched_user_name"`
	MatchedUserID          string `json:"matched_user_id"`
	MatchedUserPicture     string   `json:"matched_user_picture"`
	MatchedUserDescription string   `json:"matched_user_description"`
	MatchedUserLocation    string   `json:"matched_user_location"`
	IsOnline               bool     `json:"is_online"`
	UserInterests          []string `json:"user_interests"`
	ChatNotifications      bool     `json:"has_notification"`
}

// Returns the user's buddies who are connected
func GetBuddies(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	userBuddies, err := db.GetAllConnectedMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user connected matches:", err)
	}
	var buddies []BuddiesResponse
	var buddy BuddiesResponse
	var buddyProfile *models.ProfileInformation
	var HasNotifications bool
	var buddyID string
	for _, userMatch := range userBuddies {
		if userMatch.UserID2 == userID1 {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID1)
			buddyID = userMatch.UserID1
			if err != nil {
				log.Println("Error getting user information:", err)
			}
			fmt.Println(userID1, userMatch.UserID1)
			HasNotifications, err = db.GetUserNotifications(userID1, userMatch.UserID1)
			if err != nil {
				log.Println("Error getting HasNotifications:", err)
			}
		} else {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID2)
			buddyID = userMatch.UserID2
			if err != nil {
				log.Println("Error getting user information:", err)
			}
			fmt.Println(userID1, userMatch.UserID2)
			HasNotifications, err = db.GetUserNotifications(userID1, userMatch.UserID2)
			if err != nil {
				log.Println("Error getting HasNotifications:", err)
			}
		}
		if err != nil {
			log.Println("Error getting user information:", err)
		}
		buddy.MatchID = userMatch.ID
		buddy.Status = userMatch.Status
		buddy.MatchScore = userMatch.MatchScore
		buddy.MatchedUserName = buddyProfile.Username
		buddy.MatchedUserID = buddyID
		buddy.MatchedUserPicture = buddyProfile.Picture
		buddy.MatchedUserDescription = buddyProfile.About
		buddy.MatchedUserLocation = buddyProfile.Nation
		buddy.IsOnline = buddyProfile.IsOnline
		buddy.ChatNotifications = HasNotifications
		buddies = append(buddies, buddy)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buddies)
}

type BuddyProfile struct {
	MatchID                int      `json:"match_id"`
	MatchScore             int      `json:"match_score"`
	Status                 string   `json:"status"`
	IsOnline               bool     `json:"is_online"`
	MatchedUserName        string   `json:"matched_user_name"`
	MatchedUserPicture     string   `json:"matched_user_picture"`
	MatchedUserDescription string   `json:"matched_user_description"`
	MatchedUserLocation    string   `json:"matched_user_location"`
	UserInterests          []string `json:"user_interests"`
	UserCountry            string   `json:"user_country"`
}

// Will get the match ID  and return the buddy profile.
func GetBuddyProfile(w http.ResponseWriter, r *http.Request) {
	userID1, err := GetCurrentUserID(r)
	if err != nil {
		log.Println("Error getting user Id from token:", err)
	}
	userBuddies, err := db.GetAllConnectedMatchesByUserId(userID1)
	if err != nil {
		log.Println("Error getting user connected matches:", err)
	}
	log.Println(`userBuddies`, userBuddies)
	var buddies []BuddiesResponse
	var buddy BuddiesResponse
	var buddyProfile *models.ProfileInformation
	for _, userMatch := range userBuddies {
		if userMatch.UserID2 == userID1 {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID1)
			if err != nil {
				log.Println("Error getting user information:", err)
			}
		} else {
			buddyProfile, err = db.GetUserInformation(userMatch.UserID2)
		}
		if err != nil {
			log.Println("Error getting user information:", err)
		}
		buddy.MatchID = userMatch.ID
		buddy.Status = userMatch.Status
		buddy.MatchScore = userMatch.MatchScore
		buddy.MatchedUserName = buddyProfile.Username
		buddy.MatchedUserPicture = buddyProfile.Picture
		buddy.MatchedUserDescription = buddyProfile.About
		buddy.MatchedUserLocation = buddyProfile.Nation
		buddy.IsOnline = buddyProfile.IsOnline
		buddies = append(buddies, buddy)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(buddies)
}
