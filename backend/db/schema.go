package db

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

//Using Gorm, each table needs a specific struct to represent the data.
//Changing primary key from "ID" to "UUID" to prevent problems
type User struct {
	UUID string `gorm:"primaryKey"`
	Username string
	Password string
	Pfpfile string
	Chats []*Chat `gorm:"many2many:users_chats;"`
	CreatedAt time.Time
}

type Chat struct {
	UUID string `gorm:"primaryKey"`
	Users []*User `gorm:"many2many:users_chats;"`
	CreatedAt time.Time
}

type Message struct {
	UUID string `gorm:"primaryKey"`
	Content string
	UserID string
	ChatID string
	CreatedAt time.Time
}

func createTables(db *gorm.DB){
	//Checking if tables exist, and if they dont, create them
	hasUser := db.Migrator().HasTable(&User{})
	hasChat := db.Migrator().HasTable(&Chat{})
	hasMessage := db.Migrator().HasTable(&Message{})
	if !hasUser {
		fmt.Println("Creating table users")
		db.AutoMigrate(&User{})
		fmt.Println("Table users created")

	}
	if !hasChat {
		fmt.Println("Creating table chats")
		db.AutoMigrate(&Chat{})
		fmt.Println("Table chats created")
	}
	if !hasMessage {
		fmt.Println("Creating table messages")
		db.AutoMigrate(&Message{})
		fmt.Println("Table messages created")
	}
}