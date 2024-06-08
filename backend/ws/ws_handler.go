package ws

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"timber/backend/db"

	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Client struct {
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
        _, receivedMessage, err := c.Conn.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }

        var message *db.Message
        err = json.Unmarshal(receivedMessage, &message)
        if err != nil {
            log.Println("Error decoding JSON:", err)
            continue
        }
        db.SaveMessageToDatabase(message, database)

        c.Pool.Broadcast <- *message
    }
}

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool { return true },
}

func ServeWs(pool *Pool, w http.ResponseWriter, r *http.Request, database *gorm.DB) {
    claims, err := db.AuthToken(r, database)
    if err != 0 {
        http.Error(w, "Couldnt authenticate JWT token", err)
    }

    user, _ := db.GetUserByUsername(claims.Username, database)

    conn, err2 := Upgrade(w, r)
    if err2 != nil {
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


func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
    conn, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Println(err)
        return nil, err
    }

    return conn, nil
}

