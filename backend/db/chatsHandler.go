package db

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func CreateChat(w http.ResponseWriter, r *http.Request,db *gorm.DB){
    var invitation ChatInvitation
    log.Println("EMPEZAMOS")

    err := json.NewDecoder(r.Body).Decode(&invitation)
    if err != nil {
        http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
        return
    }

    log.Println(invitation)


    var user1 User
	var user2 User

	_ = db.Where("id = ?", invitation.Sender).First(&user1).Error;
	_ = db.Where("id = ?", invitation.Receiver).First(&user2).Error;

    log.Println(user1)
    log.Println(user2)


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

func GetChatByID(w http.ResponseWriter, r *http.Request,chatID uuid.UUID,db *gorm.DB){
    var chat Chat

    if err := db.First(&chat, "id = ?", chatID).Error; err != nil {
		http.Error(w, "Failed to get chat", http.StatusInternalServerError)
    }

    chatIDs := []uuid.UUID{chatID}
    chatWithUsers, _ := getUsersForChats( chatIDs, db)

    log.Println()
    log.Println()
    log.Println(chatIDs)
    log.Println(chatWithUsers)
    log.Println()
    log.Println()


    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(chatWithUsers); err != nil {
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

type ImageRequest struct {
    Images []string `json:"images"`
}

func ServeImage(w http.ResponseWriter, r *http.Request, imageName string) {
    filePath := filepath.Join("./uploads", fmt.Sprintf("%s.jpg", imageName))

    file, err := os.Open(filePath)
    if err != nil {
        http.Error(w, "File not found", http.StatusNotFound)
        return
    }
    defer file.Close()

    w.Header().Set("Content-Type", "image/jpg")

    _, err = io.Copy(w, file)
    if err != nil {
        http.Error(w, "Failed to serve image", http.StatusInternalServerError)
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


func CreateInvitation(w http.ResponseWriter, r *http.Request, database *gorm.DB){
    var invitation ChatInvitation

    err := json.NewDecoder(r.Body).Decode(&invitation)
    if err != nil {
        http.Error(w, `{"error": "Failed to decode JSON", "status":500"}`, http.StatusInternalServerError)
        return
    }

    var existingInvitation ChatInvitation
    err = database.Where(
        "(sender = ? AND receiver = ?) OR (sender = ? AND receiver = ?)",
        invitation.Sender, invitation.Receiver, invitation.Receiver, invitation.Sender,
    ).First(&existingInvitation).Error
    if err == nil {
        http.Error(w, `{"error": "Invitation already exists", "status": 409}`, http.StatusConflict)
        return
    }

    log.Println(existingInvitation)

    invitation.ID = uuid.New()

    err = database.Create(&invitation).Error;
    if err != nil {
        http.Error(w, `{"error": "Failed to create invitation", "status":500 "}`, http.StatusInternalServerError)
        return
    }
}

func GetInvitations(w http.ResponseWriter, r *http.Request, userID uuid.UUID, database *gorm.DB){
    var invitations []ChatInvitation
    fmt.Println(userID)

    err := database.Where("receiver = ? ",userID).Find(&invitations).Error
    if err != nil && err != gorm.ErrRecordNotFound {
        log.Printf("Error finding invitations: %v\n", err)
        http.Error(w, `{"error": "Failed to retrieve invitations"}`, http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(invitations); err != nil {
        http.Error(w, "Failed to encode chats to JSON", http.StatusInternalServerError)
        return
    }
}

func DeleteInvitation(w http.ResponseWriter, r *http.Request, invitationID uuid.UUID, database *gorm.DB){
    database.Delete(&ChatInvitation{}, invitationID)
}