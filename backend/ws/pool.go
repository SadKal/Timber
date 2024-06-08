package ws

import (
	"context"
	"log"
	"timber/backend/db"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Pool struct {
    Register   chan *Client
    Unregister chan *Client
    Clients    map[*Client]bool
    Broadcast  chan db.Message
}

func NewPool() *Pool {
    return &Pool{
        Register:   make(chan *Client),
        Unregister: make(chan *Client),
        Clients:    make(map[*Client]bool),
        Broadcast:  make(chan db.Message),
    }
}

func (pool *Pool) Start(database *gorm.DB) {
    for {
        select {
        case client := <-pool.Register:
            pool.Clients[client] = true
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
        case message := <-pool.Broadcast:
            for client := range pool.Clients {
                switch message.Type {
                case 0, 4, 6:
                    go handleNormalMessage(pool, client, message)
                case 1:
                    go handleInvitationMessage(pool, client, message)
                case 3:
                    go handleInvitationConfirmationMessage(pool, database, client, message)
                default:
                    log.Println("UNKNOWN MESSAGE TYPE:", message.Type)
                }
            }
        }
    }
}

func handleNormalMessage(pool *Pool, client *Client, message db.Message) {
    for _, chat := range client.User.Chats {
        if chat.ID == message.ChatID {
            if message.WriterUsername != client.User.Username && pool.Clients[client] {
                if err := client.Conn.WriteJSON(message); err != nil {
                    log.Println("Error writing JSON:", err)
                }
            }
        }
    }
}

func handleInvitationMessage(pool *Pool, client *Client, message db.Message) {
    receiverUUID, _ := uuid.Parse(message.Content)
    if client.User.ID == receiverUUID && pool.Clients[client] {
        if err := client.Conn.WriteJSON(message); err != nil {
            log.Println("Error writing JSON:", err)
        }
    }
}

func handleInvitationConfirmationMessage(pool *Pool,database *gorm.DB, client *Client, message db.Message) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var chat db.Chat
    if err := database.WithContext(ctx).Preload("Users").First(&chat, "id = ?", message.Content).Error; err != nil {
        log.Println("Error fetching chat:", err)
        return
    }

    for _, user := range chat.Users {
        if client.User.ID == user.ID && message.WriterUsername != client.User.Username  && pool.Clients[client] {
            client.User.Chats = append(client.User.Chats, &chat)
            if err := client.Conn.WriteJSON(message); err != nil {
                log.Println("Error writing JSON:", err)
            }
        }
    }
}

