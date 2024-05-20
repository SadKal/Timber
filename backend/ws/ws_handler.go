package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"timber/backend/db"
	"time"

	// "timber/backend/db"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

// type Handler struct {
// 	hub *Hub
// }

// func NewHandler(h *Hub) *Handler {
// 	return &Handler{
// 		hub: h,
// 	}
// }

// type CreateChatRequest struct {
// 	ID uuid.UUID `json:"id"`
// }
// func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
// 	var req CreateChatRequest

// 	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
// 		http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
//         return
//     }

// 	chat := &db.Chat{
//         ID:        uuid.New(),
//         Users:     []*db.User{},
//         CreatedAt: time.Now(),
//     }

// 	if err := database.Create(chat).Error; err != nil {
//         log.Fatalf("Failed to create chat: %v", err)
//     }

// 	h.hub.Chats[req.ID.String()] = chat

// 	w.Header().Set("Content-Type", "application/json")

//     w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(&req); err != nil {
//         http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
//         return
//     }
// }

// func (h *Handler) JoinChat(w http.ResponseWriter, r *http.Request, database *gorm.DB, params map[string]string, queryParams url.Values){
// 	conn, err := wsUpgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
//         return
// 	}

// 	claims, _ := db.AuthToken(r)

// 	chatID, _ := uuid.Parse(params["chatID"])

// 	var user db.User

// 	err = database.Where("username = ?", claims.Username).First(&user).Error
//     if err != nil {
//         http.Error(w, "Username doesnt exist exists", http.StatusNotFound)
//         return
//     }

// 	user.Connection = conn
// 	user.Message = make(chan *db.Message, 10)

// 	m := &db.Message{
// 		Content: user.Username + " has joined the chat",
// 		ChatID: chatID,
// 		Username: user.Username,
// 	}

// 	fmt.Println(m)
// }

// define a reader which will listen for
// new messages being sent to our WebSocket
// endpoint


type Client struct {
    ID   string
    Conn *websocket.Conn
    Pool *Pool
    User *db.User
}

func (c *Client) Read(database *gorm.DB) {
    defer func() {
        c.Pool.Unregister <- c
        c.Conn.Close()
    }()

    for {
        _, p, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        var message *db.Message
        err = json.Unmarshal(p, &message)
        if err != nil {
            log.Println("Error decoding JSON:", err)
            continue
        }
        message.ID = uuid.New()
        message.CreatedAt = time.Now()
        db.SaveMessageToDatabase(message, database)

        c.Pool.Broadcast <- *message
        fmt.Printf("Message Received: %+v\n", message)
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    return conn, nil
}

func ServeWs(pool *Pool, w http.ResponseWriter, r *http.Request, database *gorm.DB) {
    claims, _ := db.AuthToken(r, database)

    user, _ := db.GetUserByUsername(claims.Username, database)

    conn, err := Upgrade(w, r)
    if err != nil {
        fmt.Fprintf(w, "%+v\n", err)
    }

    client := &Client{
        Conn: conn,
        Pool: pool,
        User: user,
    }

    pool.Register <- client
    client.Read(database)
}
