package main

import (
    "github.com/joho/godotenv"
    "log"
    "os"
)

func init() {
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }
}

type NotificationType int

const (
    EmailNotification NotificationType = iota + 1
    SMSNotification
    PushNotification
)

type Notification struct {
    ID               string
    MessageContent   string
    Recipient        string
    NotificationType NotificationType
}

func sendNotification(notification *Notification) {
    log.Printf("Sending notification: %+v\n", *notification)
}

func main() {
    exampleNotification := &Notification{
        ID:               os.Getenv("NOTIFICATION_ID"),
        MessageContent:   "Hello, your order has been shipped",
        Recipient:        os.Getenv("NOTIFICATION_RECIPIENT"),
        NotificationType: EmailNotification,
    }

    sendNotification(exampleNotification)
}