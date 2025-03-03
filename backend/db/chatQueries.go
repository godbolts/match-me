package db

import (
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
)

func SaveMessage(message string, matchID int, senderID string, receiverID string) error {
	query := "INSERT INTO chat_messages (message, match_id, sender_id, receiver_id) VALUES ($1, $2, $3, $4)"
	_, err := DB.Exec(query, message, matchID, senderID, receiverID)
	if err != nil {
		return fmt.Errorf("error saving message: %v", err)
	}
	return nil
}

func SaveNotification(matchID int) error {
	query := `
		INSERT INTO unread_messages (match_id, latest_message)
		VALUES ($1, $2)
		ON CONFLICT (match_id)
		DO UPDATE SET latest_message = EXCLUDED.latest_message
	`
	currentTime := time.Now()
	_, err := DB.Exec(query, matchID, currentTime)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error saving notification status: %v", err)
	}
	return nil
}

func GetChatHistory(matchID int, offset int, limit int) ([]string, error) {
	query := "SELECT message FROM chat_messages WHERE match_id = $1 ORDER BY sent_at DESC LIMIT $2 OFFSET $3"
	var chatHistory []string
	rows, err := DB.Query(query, matchID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("error getting chat history (GetChatHistory): %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			return nil, fmt.Errorf("error scanning row (GetChatHistory): %v", err)
		}
		chatHistory = append(chatHistory, message)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over rows (GetChatHistory): %v", err)
	}
	return chatHistory, nil
}

type latestMessages struct {
	MatchID       int       `json:"match_id"`
	LatestMessage time.Time `json:"latest_message"`
}

func GetLatestMessages(matchIDs []int) ([]latestMessages, error) {
	query := `	
		SELECT match_id, latest_message
		FROM unread_messages
		WHERE match_id = ANY($1)
	`
	rows, err := DB.Query(query, pq.Array(matchIDs))
	if err != nil {
		return nil, fmt.Errorf("error fetching latest messages: %v", err)
	}
	defer rows.Close()
	var latestMessages []latestMessages
	for rows.Next() {
		var msg struct {
			MatchID       int       `json:"match_id"`
			LatestMessage time.Time `json:"latest_message"`
		}
		if err := rows.Scan(&msg.MatchID, &msg.LatestMessage); err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		latestMessages = append(latestMessages, msg)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error processing rows: %v", err)
	}
	return latestMessages, nil
}

func SaveNotifications(user1 string, user2 string, HasNotification bool) error {
	query := `
		UPDATE user_notifications 
		SET
			user_id_1_notification = CASE WHEN user_id_1 = $2 THEN $3 ELSE user_id_1_notification END,
			user_id_2_notification = CASE WHEN user_id_2 = $2 THEN $3 ELSE user_id_2_notification END
		WHERE (user_id_1 = $1 AND user_id_2 = $2)
		OR (user_id_1 = $2 AND user_id_2 = $1)
	`
	_, err := DB.Exec(query, user1, user2, HasNotification)
	if err != nil {
		log.Println(err)
		return fmt.Errorf("error saving notification status: %v", err)
	}
	return nil
}