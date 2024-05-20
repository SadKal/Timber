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
            log.Println("CONNECTED CLIENT")
            fmt.Println("Size of Connection Pool: ", len(pool.Clients))
        case client := <-pool.Unregister:
            delete(pool.Clients, client)
            log.Println("DISCONNECTED CLIENT")
        case message := <-pool.Broadcast:
            fmt.Println("Sending message to all clients in Pool")
            log.Println(message)
            for client := range pool.Clients {
                log.Println("Client: ", client.User.Username)
                if err := client.Conn.WriteJSON(message); err != nil {
                    fmt.Println(err)
                    return
                }
            }
        }
    }
}