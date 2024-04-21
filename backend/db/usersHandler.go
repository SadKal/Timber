package db

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
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

type UserRegistrationRequest struct {
    Username  string                `form:"username"`
    Password  string                `form:"password"`
    File      *multipart.FileHeader `form:"file"`
}

func RegisterUser(w http.ResponseWriter, r *http.Request, db *gorm.DB){

    // Instead of application/json im using multipartForm, because it includes an image
    // The 10 << 20 means 10^2 bytes, so 10 MB
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

    var userReq UserRegistrationRequest
    userReq.Username = r.FormValue("username")
    userReq.Password = r.FormValue("password")

    if userReq.Username == "" {
        http.Error(w, "Username is required", http.StatusBadRequest)
        return
    }
    if userReq.Password == "" {
        http.Error(w, "Password is required", http.StatusBadRequest)
        return
    }

    bytes,_ := bcrypt.GenerateFromPassword([]byte(userReq.Password), 14)

    user := User{
        ID:       uuid.New(),
        Username: userReq.Username,
        Password: string(bytes),
        CreatedAt: time.Now(),
    }

    user.Pfpfile = user.ID.String();

    file, fileHeader, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()
    userReq.File = fileHeader

    filepath := filepath.Join("./uploads/", user.ID.String())
    outFile, err := os.Create(filepath)
    if err != nil {
        http.Error(w, "Unable to create file", http.StatusInternalServerError)
        return
    }
    defer outFile.Close()

    _, err = io.Copy(outFile, file)
    if err != nil {
        http.Error(w, "Error copying file data", http.StatusInternalServerError)
        return
    }


    var existingUser User
    //Queries the db to check if the user already exists, if it doesnt, it returns an error
    err = db.Where("username = ?", user.Username).First(&existingUser).Error;
    if err == nil {
        http.Error(w, "Username already exists", http.StatusConflict)
        return
    }

    err = db.Create(&user).Error;
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(user)
    if err != nil {
        http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusCreated)
    _, err = w.Write(jsonData)
    if err != nil {
        http.Error(w, "Failed to write response", http.StatusInternalServerError)
        return
    }
}
