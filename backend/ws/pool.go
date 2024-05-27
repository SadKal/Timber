package ws

import (
	"context"
	"fmt"
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
            log.Println("CONNECTED CLIENT")
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            log.Println("DISCONNECTED CLIENT")
        case message := <-pool.Broadcast:
            // log.Println("MESSAGE", message)
            for client := range pool.Clients {
                log.Println("CURRENT CLIENT", client.User.Username)

                switch message.Type {
                case 0:
                    go handleNormalMessage(client, message)
                case 1:
                    go handleInvitationMessage(client, message)
                case 3:
                    go handleInvitationConfirmationMessage(database, client, message)
                default:
                    log.Println("UNKNOWN MESSAGE TYPE:", message.Type)
                }
            }
        }
    }
}

func handleNormalMessage(client *Client, message db.Message) {
    for _, chat := range client.User.Chats {
        log.Println("CLIENT", client.User.Username)
        log.Println("CHATS", client.User.Chats)
   

        log.Println("CURRENT CHAT ID", chat.ID)
        log.Println("CURRENT MESSAGE CHAT", message.ChatID)

        if chat.ID == message.ChatID {
            if message.WriterUsername != client.User.Username {
                if err := client.Conn.WriteJSON(message); err != nil {
                    log.Println("Error writing JSON:", err)
                }
            }
        }
    }
}

func handleInvitationMessage(client *Client, message db.Message) {
    log.Println("INVITATION: ", message.Content)
    receiverUUID, _ := uuid.Parse(message.Content)
    if client.User.ID == receiverUUID {
        if err := client.Conn.WriteJSON(message); err != nil {
            log.Println("Error writing JSON:", err)
        }
    }
}

func handleInvitationConfirmationMessage(database *gorm.DB, client *Client, message db.Message) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    var chat db.Chat
    if err := database.WithContext(ctx).Preload("Users").First(&chat, "id = ?", message.Content).Error; err != nil {
        log.Println("Error fetching chat:", err)
        return
    }

    log.Println("CHAT RECEIVED", chat.ID)
    for _, user := range chat.Users {
        if client.User.ID == user.ID && message.WriterUsername != client.User.Username {
            fmt.Println("DESTINATARIO:", client.User.Username)
            client.User.Chats = append(client.User.Chats, &chat)
            log.Println("MENSAJE A ENVIAR", message)
            if err := client.Conn.WriteJSON(message); err != nil {
                log.Println("Error writing JSON:", err)
            }
        }
    }
}
