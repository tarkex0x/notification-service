package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func NotificationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Notification sent successfully!")
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT must be set in environment variables")
	}

	http.HandleFunc("/send-notification", NotificationHandler)

	log.Printf("Starting server on port %s...", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}