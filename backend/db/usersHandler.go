package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRegistrationRequest struct {
    Username  string
    Password  string
    File      *multipart.FileHeader
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

    fmt.Println("REQUEST", userReq)

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

    uploadsDir := "./uploads"
    
    if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
        fmt.Println("Error creating uploads directory:", err)
        return
    }

    filepath := filepath.Join(uploadsDir, fmt.Sprintf("%s.jpg", user.ID.String()))
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
    token, errNum := createJWT(&user)
    if errNum == 500 {
        fmt.Println("Error while creating token")
        w.WriteHeader(errNum)
        return
    }

    fmt.Println(token)
    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusCreated)
    token.WriteTo(w)
}

func LoginUser(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
    var user User

    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, "Failed to decode JSON", http.StatusInternalServerError)
        return
    }

    if user.Username == "" {
        http.Error(w, "Username is required", http.StatusBadRequest)
        return
    }
    if user.Password == "" {
        http.Error(w, "Password is required", http.StatusBadRequest)
        return
    }

    var dbUser User
    err = db.Where("BINARY username = ?", user.Username).First(&dbUser).Error
    if err != nil {
        http.Error(w, "Username doesnt exist exists", http.StatusNotFound)
        return
    }

    err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
    if err != nil {
        http.Error(w, "Password is not correct", http.StatusUnauthorized)
        return
    }


    token, errNum := createJWT(&user)
    if errNum == 500 {
        w.WriteHeader(errNum)
        return
    }
    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusCreated)
    token.WriteTo(w)
}

func SearchUserByUsername(w http.ResponseWriter, r *http.Request, username string ,db *gorm.DB){
    var users []User

    toSearch := fmt.Sprintf("%%%s%%", username)

    db.Where("username LIKE ?", toSearch).Find(&users)

    w.Header().Set("Content-Type", "application/json")

    if err := json.NewEncoder(w).Encode(users); err != nil {
        http.Error(w, "Failed to encode chats to JSON", http.StatusInternalServerError)
        return
    }
}

func CheckAuth(w http.ResponseWriter, r *http.Request, db *gorm.DB){
    claims, err := AuthToken(r, db)
    if err != 0{
        http.Error(w, "JWT Not valid" ,http.StatusUnauthorized)
    }

    var responseBuffer bytes.Buffer

	json.NewEncoder(&responseBuffer).Encode(map[string]interface{}{
		"user": claims.Username,
        "uuid": claims.UUID,
        "expiresIn":    int(time.Until(expirationTime).Seconds()),
    })

    responseBuffer.WriteTo(w)
}

