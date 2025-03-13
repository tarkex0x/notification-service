package main

import (
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
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("DATABASE_URL is not set in .env file")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&User{}, &Preferences{}, &ContactMethod{})
	if err != nil {
		panic("failed to migrate database tables")
	}
}