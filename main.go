package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func SendNotificationResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Notification sent successfully!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		log.Fatal("PORT must be set in environment variables")
	}

	http.HandleFunc("/send-notification", SendNotificationResponse)

	log.Printf("Starting server on port %s...", serverPort)
	if err := http.ListenAndServe(":"+serverPort, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}