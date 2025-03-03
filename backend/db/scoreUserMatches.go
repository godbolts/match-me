package db

import (
	"log"
	"match_me_backend/models"
	"match_me_backend/utils"
)

func CalculateUserDistance( userID1, userID2 string) (float64, error) {
	user1, err := GetUserByID(userID1)
	if err != nil {
		log.Println("Error getting user 1", err)
	}
	user2, err := GetUserByID(userID2)
	if err != nil {
		log.Println("Error getting user 2", err)
	}
	distance := utils.GetDistanceBetweenTwoPointsOnEarth(user1.Latitude, user1.Longitude, user2.Latitude, user2.Longitude)
	if err != nil {
		log.Println("Error updating match distance", err)
	}
	matchId, err := GetMatchIdByUserIDs(userID1, userID2)
	if err != nil {
		log.Println("Error getting match id", err)
	}
	err = UpdateMatchDistance(matchId, distance)
	if err != nil {
		log.Println("Error updating match distance", err)
	}
	DistanceScoreModifier(userID1, userID2, distance)
	return distance, err
}

func CalculateMatchScore(userID1, userID2 string) (int, error) {
	user1InterestsPtr, err := GetAllUserInterest(userID1)
	if err != nil {
		log.Println("Error getting user 1 interest", err)
	}
	user1Interests := *user1InterestsPtr
	user2InterestsPtr, err := GetAllUserInterest(userID2)
	if err != nil {
		log.Println("Error getting user 2 interest", err)
	}
	user2Interests := *user2InterestsPtr
	var matchProfile []models.Interests
	for _, User1Interest := range user1Interests {
		for _, User2Interest := range user2Interests {
			if User1Interest == User2Interest {
				matchProfile = append(matchProfile, User1Interest)
			}
		}
	}
	matchScore := CalculateMatchProfile(matchProfile)
	distance, err := CalculateUserDistance(userID1, userID2)
	distanceModifier, err := DistanceScoreModifier(userID1, userID2, distance)
	matchScore = matchScore * distanceModifier
	err = UpdateUserMatchScore(userID1, userID2, matchScore)
	return matchScore, err
}

func CalculateMatchProfile(matchProfile []models.Interests) int {
	score := 0
	genreCount := 0
	playStyleCount := 0
	platformCount := 0
	communicationCount := 0
	goalsCount := 0
	sessionCount := 0
	vibeCount := 0
	languageCount := 0
	// 5 is a good score meaning there is a least one match per category
	// excluding language, platform and communication
	for _, interest := range matchProfile {
		if interest.CategoryID == GENRE {
			genreCount += 1
			score += 1
		}
		if interest.CategoryID == PLAY_STYLE {
			playStyleCount += 1
			score += 1
		}
		if interest.CategoryID == PLATFORM {
			platformCount += 1
		}
		if interest.CategoryID == COMMUNICATION {
			communicationCount += 1
		}
		if interest.CategoryID == GOALS {
			goalsCount += 1
			score += 1
		}
		if interest.CategoryID == SESSION {
			sessionCount += 1
			score += 1
		}
		if interest.CategoryID == VIBE {
			vibeCount += 1
			score += 1
		}
		if interest.CategoryID == LANGUAGE {
			languageCount += 1
		}
	}
	if languageCount < 1 || platformCount < 1 || communicationCount < 1 {
		return 0
	}
	// the score is only derived from the number of interests in the categories of Genre, PlayStyle, Goals, Session and Vibe
	return score
}

// Returns 1 if the match is valid, 0 if the match is invalid -1 if there is an error
func ValidateMatchDistancePreference(distance float64, user1Interests, userInterests2 []models.Interests) int {
	var user1DistanceLimit int
	var user2DistanceLimit int
	for _, interest := range user1Interests {
		if interest.InterestName == UP_TO_ONE_HUNDRED {
			user1DistanceLimit = 100
		}
		if interest.InterestName == ONE_HUNDRED_TO_FIVE_HUNDRED {
			user1DistanceLimit = 500
		}
		if interest.InterestName == FIVE_HUNDRED_TO_ONE_THOUSAND {
			user1DistanceLimit = 1000
		}
		if interest.InterestName == ONE_THOUSAND_AND_BEYOND {
			user1DistanceLimit = 1001
		}
	}
	for _, interest := range userInterests2 {
		if interest.InterestName == UP_TO_ONE_HUNDRED {
			user2DistanceLimit = 100
		}
		if interest.InterestName == ONE_HUNDRED_TO_FIVE_HUNDRED {
			user2DistanceLimit = 500
		}
		if interest.InterestName == FIVE_HUNDRED_TO_ONE_THOUSAND {
			user2DistanceLimit = 1000
		}
		if interest.InterestName == ONE_THOUSAND_AND_BEYOND {
			user2DistanceLimit = 1001
		}
	}
	if user1DistanceLimit == 0 || user2DistanceLimit == 0 {
		log.Println("Error: Distance limits not properly set for one or both users")
		return ERR
	}
	if float64(user1DistanceLimit) < distance || float64(user2DistanceLimit) < distance {
		return NOK
	}
	return OK
}

func DistanceScoreModifier(userID1, userID2 string, distance float64)  (int, error) {
	user1InterestsPtr, err := GetAllUserInterest(userID1)
	if err != nil {
		log.Println("Error getting user 1 interest", err)
	}
	user1Interests := *user1InterestsPtr
	user2InterestsPtr, err := GetAllUserInterest(userID2)
	if err != nil {
		log.Println("Error getting user 2 interest", err)
	}
	user2Interests := *user2InterestsPtr
	scoreModifier :=ValidateMatchDistancePreference(distance, user1Interests, user2Interests)
	return scoreModifier, err
}