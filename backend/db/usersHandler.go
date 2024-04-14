package db

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUserByUsername(w http.ResponseWriter, r *http.Request, db *gorm.DB){
	vars := mux.Vars(r)

    var user User
    if err := db.First(&user, "username = ?", vars["username"]).Error; err != nil {
        // Handle error if user is not found
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    // Marshal(format) user object into JSON format
    jsonData, err := json.Marshal(user)
    if err != nil {
        // Handle error if unable to marshal JSON
        http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

    // Set content type header to JSON
    w.Header().Set("Content-Type", "application/json")

    // Write JSON data to response writer
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(jsonData)
    if err != nil {
        // Handle error if unable to write response
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
        return
    }
}

func GetUserByUUID(w http.ResponseWriter, r *http.Request, db *gorm.DB){
	vars := mux.Vars(r)

    var user User
    if err := db.First(&user, "id = ?", vars["UUID"]).Error; err != nil {
        // Handle error if user is not found
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }
    // Marshal(format) user object into JSON format
    jsonData, err := json.Marshal(user)
    if err != nil {
        // Handle error if unable to marshal JSON
        http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

    // Set content type header to JSON
    w.Header().Set("Content-Type", "application/json")

    // Write JSON data to response writer
    w.WriteHeader(http.StatusOK)
    _, err = w.Write(jsonData)
    if err != nil {
        // Handle error if unable to write response
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
        return
    }
}