package main

import (
	"net/http"
	"timber/backend/db"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func setupRoutes(database *gorm.DB) {
	router := mux.NewRouter()

	router.Use(corsMiddleware)
	//Call the router with a specific function to be able to pass the database as a parameter
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        db.GetUserByUsername(w, r, database)
    })

	router.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
        db.RegisterUser(w, r, database)
    })

	router.HandleFunc("/users/{UUID}", func(w http.ResponseWriter, r *http.Request) {
        db.GetUserByUUID(w, r, database)
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
		if r.Method == http.MethodOptions {
            // Respond with the required CORS headers
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            w.WriteHeader(http.StatusOK)
            return
        }

        // Allow requests from any origin on localhost
        w.Header().Set("Access-Control-Allow-Origin", "*")

        // Proceed to the next middleware or handler
        next.ServeHTTP(w, r)
    })
}
