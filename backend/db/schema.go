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
    ID        uuid.UUID 		`gorm:"primaryKey;type:char(36)" json:"id"`
    Username  	string    		`gorm:"uniqueIndex" json:"username"`
    Password  	string    		`json:"password"`
    Pfpfile   	string    		`json:"pfpfile"`
    Chats     	[]*Chat   		`gorm:"many2many:users_chats;" json:"-"`
    CreatedAt 	time.Time 		`json:"created_at"`
}

type Chat struct {
    ID        uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
    Users     []*User   `gorm:"many2many:users_chats;" json:"-"`
    CreatedAt time.Time `json:"created_at"`
}

type Message struct {
    ID        uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
	Type 	  int    	`gorm:"type:tinyint" json:"type"`
    Content   string    `json:"content"`
	WriterUsername string		`json:"username"`
    UserID  uuid.UUID 	`json:"user_id"`
    User      User      `gorm:"foreignKey:UserID" json:"-"`
    ChatID    uuid.UUID `json:"chat_id"`
    Chat      Chat      `gorm:"foreignKey:ChatID" json:"-"`
    CreatedAt time.Time `json:"created_at"`
}

type ChatInvitation struct {
    ID uuid.UUID `gorm:"primaryKey;type:char(36)" json:"id"`
    SenderUsername  string 	`json:"senderUsername"`
	Sender User `gorm:"foreignKey:SenderUsername" json:"-"`
	ReceiverUsername string `json:"receiverUsername"`
	Receiver User `gorm:"foreignKey:ReceiverUsername" json:"-"`
	ChatID uuid.UUID `json:"chat_id"`
	Chat Chat `gorm:"foreignKey:ChatID;references:ID" json:"-"`
}

func GetUserByUsername(username string, database *gorm.DB) (*User, error) {
	var user User

	if err := database.First(&user, "username=?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
}


func createTables(db *gorm.DB){
	//Checking if tables exist, and if they dont, create them
	hasUser := db.Migrator().HasTable(&User{})
	hasChat := db.Migrator().HasTable(&Chat{})
	hasMessage := db.Migrator().HasTable(&Message{})
	hasChatInvitation := db.Migrator().HasTable(&ChatInvitation{})

	// fmt.Println(hasUser)
	// fmt.Println(hasChat)
	// fmt.Println(hasMessage)
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
	if !hasChatInvitation {
		fmt.Println("Creating table chat invitations")
		db.AutoMigrate(&ChatInvitation{})
		fmt.Println("Table chat invitations created")
	}

	//Create chat example
	/*var user1 User
	var user2 User

	_ = db.Where("username = ?", "Test").First(&user1).Error;
	_ = db.Where("username = ?", "Test2").First(&user2).Error;

	chat := &Chat{
        ID:        uuid.New(),
        Users:     []*User{&user1, &user2},
        CreatedAt: time.Now(),
    }

	if err := db.Create(chat).Error; err != nil {
        log.Fatalf("Failed to create chat: %v", err)
    }*/


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
	// var user User
	// db.First(&user, "username = ?", "Pepe")
	// var chat Chat
	// db.First(&chat)
	// fmt.Println(chat.ID);
	// message := &Message{
	// 	ID: uuid.New(),
	// 	Content: "Bofofofof",
	// 	UserID: user.ID,
	// 	ChatID: chat.ID,
	// }
	// if db.Create(&message).Error != nil {
	// 	log.Panic("Unable to create message.")
	// }
}