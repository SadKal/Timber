package main

import (
	// "fmt"
    "timber/backend/db"

	// "github.com/google/uuid"
)

// "fmt"
// "log"
// "net/http"
// "log"
// "time"
// "github.com/gorilla/websocket"

// var upgrader = websocket.Upgrader{
//     ReadBufferSize:  1024,
//     WriteBufferSize: 1024,
// }

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Home page")
// }

// func reader(conn *websocket.Conn){
// 	for {
// 		messageType, p, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}
// 		fmt.Println(string(p))

// 		if err := conn.WriteMessage(messageType, p); err != nil {
// 			log.Println(err)
// 			return
// 		}
// 	}
// }

// func wsEndpoint(w http.ResponseWriter, r *http.Request) {
// 	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

// 	ws, err := upgrader.Upgrade(w, r, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	ticker := time.NewTicker(5 * time.Second)
// 	quit := make(chan struct{})
// 	go func() {
// 		for {
// 			select {
// 			case <- ticker.C:
// 				ws.WriteMessage(1, []byte("Test"))
// 			case <- quit:
// 				ticker.Stop()
// 				return
// 			}
// 		}
// 	}()

// 	log.Println("Client connected")
// 	err = ws.WriteMessage(1, []byte("Hi Client!"))
//     if err != nil {
//         log.Println(err)
//     }

// 	reader(ws)
// }

// func setupRoutes(){
// 	http.HandleFunc("/", homePage)
// 	http.HandleFunc("/ws", wsEndpoint)
// }


func main(){
	// fmt.Println("Hello, world!")
	// setupRoutes()
	// log.Fatal(http.ListenAndServe(":8080", nil))
    // id := uuid.New()
    // fmt.Println(id.String())
    db.Connect()
}