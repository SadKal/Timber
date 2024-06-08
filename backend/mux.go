package main

import (
	"net/http"
	"path/filepath"
	"timber/backend/db"
	"timber/backend/ws"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	"gorm.io/gorm"
)


func setupRoutes(database *gorm.DB, router *mux.Router) {
	router.Use(corsMiddleware)

    pool := ws.NewPool()
    go pool.Start(database)

    //AUTH
	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        db.RegisterUser(w, r, database)
    }).Methods("POST", "OPTIONS")

    router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        db.LoginUser(w, r, database)
    }).Methods("POST", "OPTIONS")

    router.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
        db.CheckAuth(w, r, database)
    }).Methods("GET", "OPTIONS")


    //CHATS
    router.HandleFunc("/chats", func(w http.ResponseWriter, r *http.Request) {
        db.GetChatsFromUser(w, r, database)
    }).Methods("POST", "OPTIONS")

    router.HandleFunc("/chats/{chatID}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        chatID, _ := uuid.Parse(vars["chatID"])
        db.GetChatByID(w, r, chatID, database)
    }).Methods("GET", "OPTIONS")

    router.HandleFunc("/createchat", func(w http.ResponseWriter, r *http.Request) {
        db.CreateChat(w, r, database)
    }).Methods("POST", "OPTIONS")


    //INVITATIONS
    router.HandleFunc("/invitations", func(w http.ResponseWriter, r *http.Request) {
        db.CreateInvitation(w, r, database)
    }).Methods("POST", "OPTIONS")

    router.HandleFunc("/invitations/{user_uuid}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        user_uuid, _ := uuid.Parse(vars["user_uuid"])
        db.GetInvitations(w, r, user_uuid ,database)
    }).Methods("GET", "OPTIONS")

    router.HandleFunc("/invitations/{invitation_uuid}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        invitation_uuid, _ := uuid.Parse(vars["invitation_uuid"])
        db.DeleteInvitation(w, r, invitation_uuid ,database)
    }).Methods("DELETE", "OPTIONS")


    //MESSAGES
    router.HandleFunc("/messages/{chatID}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        chatID, _ := uuid.Parse(vars["chatID"])
        db.GetMessagesForChat(w, r, chatID, database)
    }).Methods("GET", "OPTIONS")

    router.HandleFunc("/messages/{messageID}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        messageID, _ := uuid.Parse(vars["messageID"])
        db.DeleteMessage(w, r, messageID, database)
    }).Methods("DELETE", "OPTIONS")

    router.HandleFunc("/messages/{messageID}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        messageID, _ := uuid.Parse(vars["messageID"])
        db.EditMessage(w, r, messageID, database)
    }).Methods("PUT", "OPTIONS")


    //USERS
    router.HandleFunc("/users/{username}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        username := vars["username"]
        db.SearchUserByUsername(w, r, username, database)
    }).Methods("GET", "OPTIONS")


    //IMAGES
    router.HandleFunc("/images/{image}", func(w http.ResponseWriter, r *http.Request) {
        vars := mux.Vars(r)
        imageName := filepath.Base(vars["image"])
        db.ServeImage(w, r, imageName)
    }).Methods("GET", "OPTIONS")

    //WEBSOCKET
    router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
        ws.ServeWs(pool, w, r, database)
    })

    http.Handle("/", router)
}


func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		/*Este lo pongo en español porque que liada el cors.
		Basicamente, antes de enviar los datos el cliente hace una llamada
		para comprobar que tiene permitido conectarse con el servidor
		Esto, por lo general, parece estar bloqueado, por lo que se debe permitir el
		acceso de "cors". He puesto que permita todos los origenes, ya que va a estar
		corriendo en local*/
        origin := r.Header.Get("Origin")
		if r.Method == http.MethodOptions {
            // Respond with the required CORS headers
            w.Header().Set("Access-Control-Allow-Origin", origin)
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
            w.Header().Set("Access-Control-Allow-Credentials", "true")
            w.WriteHeader(http.StatusOK)
            return
        }

        // Allow requests from any origin on localhost
        w.Header().Set("Access-Control-Allow-Origin", origin)
        w.Header().Set("Access-Control-Allow-Credentials", "true")

        // Proceed to the next middleware or handler
        next.ServeHTTP(w, r)
    })
}