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
		errMsg := method + " method is allowed"
		http.Error(w, errMsg, http.StatusMethodNotAllowed)
		log.Printf("Method not allowed: %s, required: %s\n", r.Method, method)
		return false
	}
	return true
}

func decodeJSON(w http.ResponseWriter, r *http.Request, target interface{}) bool {
	if err := json.NewDecoder(r.Body).Decode(target); err != nil {
		errMsg := "Error decoding JSON: " + err.Error()
		http.Error(w, errMsg, http.StatusBadRequest)
		log.Printf("Failed to decode JSON: %v\n", err)
		return false
	}
	return true
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		// If encoding fails, log the error and don't expose it to the client
		log.Printf("Failed to encode JSON: %v\n", err)
		http.Error(w, "Failed to process request", http.StatusInternalServerError)
	}
}