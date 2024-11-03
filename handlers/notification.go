package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Notification struct {
	ID      string `json:"id"`
	Message string `json:"message"`
}

var notifications []Notification

func main() {
	http.HandleFunc("/create", createNotificationHandler)
	http.HandleFunc("/send", sendNotificationHandler)

	startServer()
}

func startServer() {
	port := getServerPort()
	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}

func createNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if ensureMethod(w, r, "POST") {
		var newNotification Notification
		if decodeJSON(w, r, &newNotification) {
			notifications = append(notifications, newNotification)
			respondWithJSON(w, http.StatusCreated, newNotification)
		}
	}
}

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if ensureMethod(w, r, "GET") {
		respondWithJSON(w, http.StatusOK, notifications)
	}
}

func ensureMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, method+" method is allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func decodeJSON(w http.ResponseWriter, r *http.Request, target interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}