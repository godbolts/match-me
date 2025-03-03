package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"match_me_backend/db"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/websocket"
)

var connections = make(map[string]*websocket.Conn)
var mu sync.Mutex

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("Origin") == "http://localhost:3000"
	},
}

type MessageData struct {
	SenderID   string `json:"senderID"`
	ReceiverID string `json:"receiverID"`
	Message    string `json:"message"`
	Type       string `json:"type"`
	Username   string `json:"username"`
}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to upgrade to WebSocket:", err)
		return
	}
	defer conn.Close()
	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Println("No userID provided")
		return
	}
	mu.Lock()
	connections[userID] = conn
	mu.Unlock()
	log.Printf("WebSocket connection established for userID: %s\n", userID)
	log.Printf("Number of connections: %d\n", len(connections))
	log.Printf("Connections: %d\n", connections)
	defer func() {
		mu.Lock()
		delete(connections, userID)
		mu.Unlock()
		log.Printf("Websocket connection closed for userID: %s\n", userID)
		conn.Close()
	}()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}
		var msgData MessageData
		if err := json.Unmarshal(message, &msgData); err != nil {
			log.Println("Error unmarshaling message:", err)
			continue
		}
		if msgData.Type == "login" || msgData.Type == "logout" {
			mu.Lock()
			for _, client := range connections {
				err := client.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("Error sending login notification:", err)
				}
			}
			mu.Unlock()
		}
		mu.Lock()
		senderConn, senderOnline := connections[msgData.SenderID]
		receiverConn, receiverOnline := connections[msgData.ReceiverID]
		mu.Unlock()
		if msgData.Type == "typing" || msgData.Type == "stopTyping" {
			if receiverOnline {
				err := receiverConn.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Println("Error sending typing status to receiver:", err)
				}
			}
		}
		if senderOnline {
			err := senderConn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error sending message to sender:", err)
			}
		}
		if receiverOnline {
			err := receiverConn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				log.Println("Error sending message to receiver:", err)
			}
		}
	}
}

func ChatDataHandler(w http.ResponseWriter, r *http.Request) {
	senderID := r.URL.Query().Get("senderID")
	matchID := r.URL.Query().Get("matchID")
	if senderID == "" || matchID == "" {
		http.Error(w, "Both senderID and matchID are required", http.StatusBadRequest)
		return
	}
	receiverID, err := db.GetReceiverID(matchID, senderID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error fetching receiver ID: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(receiverID)
}

func SaveMessageHandler(w http.ResponseWriter, r *http.Request) {
	var messageData struct {
		MatchID    int    `json:"matchID"`
		SenderID   string `json:"senderID"`
		ReceiverID string `json:"receiverID"`
		Message    string `json:"message"`
	}
	err := json.NewDecoder(r.Body).Decode(&messageData)
	if err != nil {
		log.Printf("ERROR: Failed to save message to database. Error: %v, Arguments: %v", err, messageData)
		http.Error(w, "Error parsing request body", http.StatusBadRequest)
		return
	}
	err = db.SaveMessage(messageData.Message, messageData.MatchID, messageData.SenderID, messageData.ReceiverID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving message: %v", err), http.StatusInternalServerError)
		return
	}
	err = db.SaveNotification(messageData.MatchID)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving notification: %v", err), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Message saved successfully")
}

func ChatHistoryHandler(w http.ResponseWriter, r *http.Request) {
	matchIDStr := r.URL.Query().Get("matchID")
	if matchIDStr == "" {
		http.Error(w, "Missing matchID", http.StatusBadRequest)
		return
	}
	matchID, err := strconv.Atoi(matchIDStr)
	if err != nil {
		http.Error(w, "Invalid matchID", http.StatusBadRequest)
	}
	offsetStr := r.URL.Query().Get("offset")
	offset := 0
	if offsetStr != "" {
		offset, err = strconv.Atoi(offsetStr)
		if err != nil {
			http.Error(w, "Invalid offset", http.StatusBadRequest)
			return
		}
	}
	limit := 15
	chatHistory, err := db.GetChatHistory(matchID, offset, limit)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting chat history: %v", err), http.StatusInternalServerError)
		return
	}
	for i, j := 0, len(chatHistory)-1; i < j; i, j = i+1, j-1 {
		chatHistory[i], chatHistory[j] = chatHistory[j], chatHistory[i]
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chatHistory)
}

type LatestMessageRequest struct {
	MatchIDs []int `json:"match_ids"`
}

func ChatMessageHandler(w http.ResponseWriter, r *http.Request) {
	var request LatestMessageRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	matchIDs := request.MatchIDs
	latestMessages, err := db.GetLatestMessages(matchIDs)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error getting latest message info: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(latestMessages)
}

type SetNotificationRequest struct {
	User1           string `json:"user1"`
	User2           string `json:"user2"`
	HasNotification bool   `json:"has_notification"`
}

func ChatNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var request SetNotificationRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	User1 := request.User1
	User2 := request.User2
	HasNotification := request.HasNotification
	err = db.SaveNotifications(User1, User2, HasNotification)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error saving notifications: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Notification saved!")
}