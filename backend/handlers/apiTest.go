package handlers

import (
	"encoding/json"

	"net/http"
)

//to test any GET function use postman and run localhost:4000/test
func GetTestResultHandler(w http.ResponseWriter, r *http.Request) {
	successMessage := "No test is set!"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successMessage)
}
