package db

import (
	"fmt"
	"log"
	"match_me_backend/auth"
	"strconv"
	"time"

	"golang.org/x/exp/rand"
	
)


/*
Run only once, to create profiles for demo users.

To re spawn all demo users entries from the tables should be removed first from the following tables.

TRUNCATE TABLE user_interests;
TRUNCATE TABLE users;
TRUNCATE TABLE user_matches;
TRUNCATE TABLE profiles

*/

// RUN ONLY ONCE WHEN NEEDED. DO NOT RUN AGAIN BEFORE REMOVING ALL ENTRIES FROM THE TABLES
func InitDemoUsers() bool {
	for i := 0; i < DEMO_USER_COUNT; i++ {
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		hashedPassword, err := auth.HashPassword(iStr)
		if err != nil {
			log.Println("Error hashing password: ", err)
		}
		err = SaveUser(email, hashedPassword)
		if err != nil {
			log.Println("Error saving user: ", err)
		}
	}
	CreateProfile()
	log.Println("Demo users initialized")
	return true
}

func CreateProfile() {
	fmt.Println("Creating profiles")
	rand.Seed(uint64(time.Now().UnixNano()))
	birthdate, err := time.Parse("2006-01-02", "1999-01-01")
	var latitude float64
	var longitude float64
	imageIndex := 0
	distanceInterest := 49
	if err != nil {
		log.Println("Error parsing birthdate: ", err)
	}
	for i := 0; i < DEMO_USER_COUNT; i++ {
		if imageIndex == 13 {
			imageIndex = 0
		}
		if distanceInterest == 53 {
			distanceInterest = 49
		}
		iStr := strconv.Itoa(i)
		email := iStr + "@" + iStr + ".com"
		uuid, err := GetUserUUIDFromUserEmail(email) //
		if err != nil {
			log.Println("Error getting user uuid: ", err)
		}
		// tartu is base
		latitude = 58.378025 + float64(i)
		longitude = 26.728493 + float64(i)
		err = SetUsername(uuid, "User"+iStr)
		if err != nil {
			log.Println("Error setting username: ", err)
		}
		err = SetBirthdate(uuid, birthdate) // 1999-01-01 all user have the same birthdate
		if err != nil {
			log.Println("Error setting birthdate: ", err)
		}
		err = SetAbout(uuid, "I am a user "+iStr)
		if err != nil {
			log.Println("Error setting about: ", err)
		}
		// create a picture path
		picturePath := fmt.Sprintf("bot%d.png", imageIndex)
		err = SetPicturePath(uuid, picturePath)
		imageIndex++
		if err != nil {
			log.Println("Error setting picture path: ", err)
			
		}
		
		// tartu is base
		latitude = 58.378025 + float64(i)
		longitude = 26.728493 + float64(i)
		err = SetCity(uuid, "Estonia", "Tartu County", "Tartu", latitude, longitude) // all users are from Tartu random lat and long just stars adding distance to users
		if err != nil {
			log.Println("Error setting city: ", err)
		}
		// Add genre
		for j := 0; j <= 3; j++ {
			rndNum, err := GenerateRandomNumber(1, 10)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndNum, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add play style
		for k := 0; k <= 3; k++ {
			rndNum, err := GenerateRandomNumber(11, 16)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndNum, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add platform
		// rndPlatform, err := GenerateRandomNumber(17, 18)
		err = AddInterestToUser(17, uuid)
		if err != nil {
			log.Printf("Error adding interest to user %s: %v", uuid, err)
		}
		// Add communication
		for m := 0; m <= 2; m++ {
			// rndNum, err := GenerateRandomNumber(22, 23)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(22, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add goals
		for n := 0; n <= 2; n++ {
			rndNum, err := GenerateRandomNumber(27, 31)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndNum, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add session
		rndSession, err := GenerateRandomNumber(32, 34)
		err = AddInterestToUser(rndSession, uuid)
		if err != nil {
			log.Printf("Error adding interest to user %s: %v", uuid, err)
		}
		// Add vibe
		for n := 0; n <= 3; n++ {
			rndVibe, err := GenerateRandomNumber(35, 41)
			if err != nil {
				log.Println("Error generating random number: ", err)
			}
			err = AddInterestToUser(rndVibe, uuid)
			if err != nil {
				log.Printf("Error adding interest to user %s: %v", uuid, err)
			}
		}
		// Add language
		// rndLanguage, err := GenerateRandomNumber(42, 45)
		err = AddInterestToUser(42, uuid)
		if err != nil {
			log.Printf("Error adding Language interest to user %s: %v", uuid, err)
		}
		err = AddInterestToUser(distanceInterest, uuid)
		distanceInterest++
	}
}

func GenerateRandomNumber(min, max int) (int, error) {
	if min > max {
		return 0, fmt.Errorf("invalid range: min (%d) cannot be greater than max (%d)", min, max)
	}
	return rand.Intn(max-min+1) + min, nil
}
