package ws

import (
	"fmt"
	"log"
	"timber/backend/db"
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

func (pool *Pool) Start() {
    for {
        select {
        case client := <-pool.Register:
            pool.Clients[client] = true
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
            for client := range pool.Clients {
                message := db.NewMessage(0, client.User.Username + " connected", client.User.Username, "3")
                fmt.Println("CLIENTE", client.User.ID)
                client.Conn.WriteJSON(message)
            }
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
            for client := range pool.Clients {
                message := db.NewMessage(0, client.User.Username + " disconnected", client.User.Username, "3")
                client.Conn.WriteJSON(message)
            }
        case message := <-pool.Broadcast:
            fmt.Println("Sending message to all clients in Pool")
            log.Println(message)
            for client := range pool.Clients {
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
        }
    }
}