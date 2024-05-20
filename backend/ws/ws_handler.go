package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"timber/backend/db"

	// "timber/backend/db"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

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
        // message.CreatedAt = time.Now()
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
