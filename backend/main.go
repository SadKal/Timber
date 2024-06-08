package main

import (
	"log"
	"net/http"
	"timber/backend/db"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main(){
    err := godotenv.Load(".env")
    if err != nil{
        log.Fatalf("Error loading .env file: %s", err)
    }

    db := db.Connect()
	router := mux.NewRouter()
    setupRoutes(db, router)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal("HTTP server error: ", err)
	}
}