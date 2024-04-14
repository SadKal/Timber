package main

import (
	"net/http"
	"timber/backend/db"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func setupRoutes(database *gorm.DB) {
	router := mux.NewRouter()
	// router.HandleFunc("/", HomeHandler)

	//Call the router with a specific function to be able to pass the database as a parameter
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
        db.GetUserByUsername(w, r, database)
    })
	router.HandleFunc("/users/{UUID}", func(w http.ResponseWriter, r *http.Request) {
        db.GetUserByUUID(w, r, database)
    })


	http.Handle("/", router)
}