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
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func createNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	var newNotification Notification
	err := json.NewDecoder(r.Body).Decode(&newNotification)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	notifications = append(notifications, newNotification)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newNotification)
}

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(notifications)
}