package db

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func CreateChat(w http.ResponseWriter, r *http.Request,db *gorm.DB){
    var user1 User
	var user2 User

	_ = db.Where("username = ?", "Goku").First(&user1).Error;
	_ = db.Where("username = ?", "Test").First(&user2).Error;

	chat := &Chat{
        ID:        uuid.New(),
        Users:     []*User{&user1, &user2},
        CreatedAt: time.Now(),
    }

	if err := db.Create(chat).Error; err != nil {
        log.Fatalf("Failed to create chat: %v", err)
    }

	w.Header().Set("Content-Type", "application/json")
    response := map[string]uuid.UUID{"id": chat.ID}
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
    }
}

type requestUUID struct {
	Uuid string `json:"uuid"`
}

type ChatWithUsers struct {
    ChatID uuid.UUID `json:"chat_id"`
    Users  []*User   `json:"users"`
}

func GetChatsFromUser(w http.ResponseWriter, r *http.Request,db *gorm.DB){
	var request requestUUID
	var chatIDs []uuid.UUID

	err := json.NewDecoder(r.Body).Decode(&request)
    if err != nil {
        http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
        return
    }

	var user User

	if err := db.Preload("Chats").First(&user, "id = ?", request.Uuid).Error; err != nil {
		http.Error(w, "Failed to get user chats", http.StatusInternalServerError)
    }
	for i := 0; i < len(user.Chats); i++ {
		chatIDs = append(chatIDs, user.Chats[i].ID)
	}

	chats, _ := getUsersForChats(chatIDs, db)

	w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(chats); err != nil {
        http.Error(w, "Failed to encode chats to JSON", http.StatusInternalServerError)
        return
    }
}

func getUsersForChats(chatIDs []uuid.UUID, db *gorm.DB) ([]ChatWithUsers, error) {
    var chatsWithUsers []ChatWithUsers

    for _, chatID := range chatIDs {
        var chat Chat
        if err := db.Preload("Users").First(&chat, "id = ?", chatID).Error; err != nil {
            // Handle error if chat is not found
            return nil, err
        }

        // Append chat data along with its users to the result
        chatsWithUsers = append(chatsWithUsers, ChatWithUsers{
            ChatID: chatID,
            Users:  chat.Users,
        })
    }

    return chatsWithUsers, nil
}

func GetMessagesForChat(w http.ResponseWriter, r *http.Request, chatID uuid.UUID, db *gorm.DB) {
    var messages []Message
    if err := db.Order("created_at desc").Where("chat_id = ?", chatID).Find(&messages).Error; err != nil {
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(messages); err != nil {
        http.Error(w, "Failed to encode chats to JSON", http.StatusInternalServerError)
        return
    }
}
func SaveMessageToDatabase(message *Message, database *gorm.DB){
    err := database.Create(&message).Error;
    if err != nil {
        log.Println("Failed to save message to database")
        return
    }
}