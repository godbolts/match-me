package routes

import (
	"match_me_backend/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	fileDirectory := "../frontend/src/components/Assets/ProfilePictures"

	router.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir(fileDirectory))))
	// user routes
	router.HandleFunc("/authorization", handlers.AuthorizationHandler).Methods("POST")
	router.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/login/api", handlers.LoginAPIHandler).Methods("GET")
	router.HandleFunc("/logout", handlers.LogoutHandler).Methods("GET")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	router.HandleFunc("/users/{id}", handlers.GetUserHandler).Methods("GET")
	router.HandleFunc("/users/{id}/profile", handlers.GetUserProfileHandler).Methods("GET")
	router.HandleFunc("/users/{id}/bio", handlers.GetUserBioHandler).Methods("GET")
	router.HandleFunc("/me", handlers.GetLightCurrentUserHandler).Methods("GET")
	router.HandleFunc("/me/profile", handlers.GetCurrentUserHandler).Methods("GET")
	router.HandleFunc("/me/bio", handlers.GetMeBioHandler).Methods("GET")
	router.HandleFunc("/me/uuid", handlers.GetCurrentUserUUID).Methods("GET")

	router.HandleFunc("/online", handlers.GetOnlineStatus).Methods("GET")
	router.HandleFunc("/online/{id}", handlers.GetOtherOnlineStatus).Methods("GET")

	// Demo and Test routes
	router.HandleFunc("/test", handlers.GetTestResultHandler).Methods("GET")
	router.HandleFunc("/spawn/bots", handlers.GetDemoUsers).Methods("GET")

	// Profile routes
	router.HandleFunc("/userInterests", handlers.GetUserInterests).Methods("GET")
	router.HandleFunc("/userInterest", handlers.UpdateUserInterest).Methods("POST")
	router.HandleFunc("/interests/{id}", handlers.GetUserBioHandler).Methods("GET")
	router.HandleFunc("/interests", handlers.GetInterests).Methods("GET")
	router.HandleFunc("/username", handlers.PostUsername).Methods("POST")
	router.HandleFunc("/city", handlers.PostCity).Methods("POST")
	router.HandleFunc("/about", handlers.PostAbout).Methods("POST")
	router.HandleFunc("/birthdate", handlers.PostBirthdate).Methods("POST")
	router.HandleFunc("/picture", handlers.PostProfilePictureHandler).Methods("POST")
	router.HandleFunc("/picture/remove", handlers.PostProfileRPictureRemoveHandler).Methods("POST")

	router.HandleFunc("/browserlocation", handlers.BrowserHandler).Methods("POST")

	// match routes
	router.HandleFunc("/matches", handlers.GetMatches).Methods("GET")
	router.HandleFunc("/requests", handlers.GetRequests).Methods("GET")
	router.HandleFunc("/matches/request", handlers.RequestMatch).Methods("PUT")
	router.HandleFunc("/matches/connect", handlers.ConfirmMatch).Methods("PUT")
	router.HandleFunc("/matches/block", handlers.BlockMatch).Methods("PUT")
	router.HandleFunc("/matches/remove", handlers.RemoveMatch).Methods("PUT")
	router.HandleFunc("/connections", handlers.GetConnections).Methods("GET")
	router.HandleFunc("/buddies", handlers.GetBuddies).Methods("GET")
	router.HandleFunc("/recommendations", handlers.GetRecommendationsHandler).Methods("GET")

	// chat routes
	router.HandleFunc("/ws", handlers.WebsocketHandler)
	router.HandleFunc("/receiver", handlers.ChatDataHandler).Methods("GET")
	router.HandleFunc("/saveMessage", handlers.SaveMessageHandler).Methods("POST")
	router.HandleFunc("/chatHistory", handlers.ChatHistoryHandler).Methods("GET")
	router.HandleFunc("/latestMessage", handlers.ChatMessageHandler).Methods("POST")
	router.HandleFunc("/saveNotification", handlers.ChatNotificationHandler).Methods("POST")

	return router
}
