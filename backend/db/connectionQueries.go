package db

import (
	"fmt"
	"log"
)

func SetFirstConnection(user_id1, user_id2 string) error {
	userQuery := "INSERT INTO Connections (UserID1, UserID2) VALUES ($1, $2)"
	_, err := DB.Exec(userQuery, user_id1, user_id2)
	if err != nil {
		log.Printf("Error connecting username for uuid=%s: %v", user_id1, err)
		return fmt.Errorf("could not connect username: %w", err)
	}
	return nil
}

func SetAccepted(user_id1, user_id2 string) error {
	userQuery := "UPDATE Connections SET Status = 'accepted' WHERE UserID1 = $1 AND UserID2 = $2"
	_, err := DB.Exec(userQuery, user_id2, user_id1)
	if err != nil {
		log.Printf("Error accepting username for uuid=%s: %v", user_id1, err)
		return fmt.Errorf("could not accept username: %w", err)
	}
	return nil
}

func Setblock(user_id1, user_id2 string) error {
	userQuery := "UPDATE Connections SET Status = 'blocked' WHERE (UserID1 = $1 AND UserID2 = $2) OR (UserID1 = $2 AND UserID2 = $1)"
	_, err := DB.Exec(userQuery, user_id1, user_id2)
	if err != nil {
		log.Printf("Error blocking username for uuid=%s: %v", user_id1, err)
		return fmt.Errorf("could not block username: %w", err)
	}
	return nil
}
