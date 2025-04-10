package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"size:255"`
	Email          string `gorm:"unique;not null"`
	PhoneNumber    string `gorm:"size:100"`
	Preferences    Preferences
	ContactMethods []ContactMethod `gorm:"foreignKey:UserID"`
}

type Preferences struct {
	gorm.Model
	UserID      uint
	DarkMode    bool `json:"dark_mode"`
	EmailNotify bool `json:"email_notify"`
}

type ContactMethod struct {
	gorm.Model
	UserID    uint
	Type      string `gorm:"size:50"`
	Value     string `gorm:"size:255"`
	IsPrimary bool
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL is not set in .env file")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %s", err.Error())
	}

	err = db.AutoMigrate(&User{}, &Preferences{}, &ContactMethod{})
	if err != nil {
		log.Fatalf("Failed to migrate database tables: %s", err.Error())
	}

	fmt.Println("Database connection and migration successful")
}
