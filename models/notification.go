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
    sendNotificationBasedOnType(notification)
}

func sendNotificationBasedOnType(notification *Notification) {
    switch notification.NotificationType {
    case EmailNotification:
        sendEmailNotification(notification)
    case SMSNotification:
        sendSMSNotification(notification)
    case PushNotification:
        sendPushNotification(notification)
    default:
        log.Println("Invalid notification type")
    }
}

func sendEmailNotification(notification *Notification) {
    log.Printf("Sending email notification to %s: %+v\n", notification.Recipient, *notification)
}

func sendSMSNotification(notification *Notification) {
    log.Printf("Sending SMS notification to %s: %+v\n", notification.Recipient, *notification)
}

func sendPushNotification(notification *Notification) {
    log.Printf("Sending push notification to %s: %+v\n", notification.Recipient, *notification)
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