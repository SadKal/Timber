package db

import (
	"fmt"
	// "log"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

//Using Gorm, each table needs a specific struct to represent the data.
//Changing primary key from "ID" to "UUID" to prevent problems
type User struct {
    ID        uuid.UUID `gorm:"primaryKey;type:char(36)"`
    Username  string
    Password  string
    Pfpfile   string
    Chats     []*Chat `gorm:"many2many:users_chats;"`
    CreatedAt time.Time
}

type Chat struct {
    ID        uuid.UUID `gorm:"primaryKey;type:char(36)"`
    Users     []*User `gorm:"many2many:users_chats;"`
    CreatedAt time.Time
}

type Message struct {
    ID        uuid.UUID `gorm:"primaryKey;type:char(36)"`
    Content   string
    UserID    uuid.UUID
    User      User `gorm:"foreignKey:UserID;references:ID"`
    ChatID    uuid.UUID
    Chat      Chat `gorm:"foreignKey:ChatID;references:ID"`
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

	// user := &User{
	// 	ID: uuid.New(),
	// 	Username: "Pepe",
	// 	Password: "Pepe",
	// 	Pfpfile: ".",
	// }
	// if db.Create(&user).Error != nil {
	// 	log.Panic("Unable to create user.")
	// }
	// chat := &Chat{
	// 	ID: uuid.New(),
	// }
	// if db.Create(&chat).Error != nil {
	// 	log.Panic("Unable to create user.")
	// }
	// message := &Message{
	// 	ID: uuid.New(),
	// 	Content: "Bofofofof",
	// }
	// if db.Create(&message).Error != nil {
	// 	log.Panic("Unable to create user.")
	// }
}