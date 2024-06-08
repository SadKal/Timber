package db

import (
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


func RegisterUser(w http.ResponseWriter, r *http.Request, database *gorm.DB){
    // Instead of application/json im using multipartForm, because it includes an image
    // The 10 << 20 means 10^2 bytes, so 10 MB
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Unable to parse form", http.StatusBadRequest)
        return
    }

	var userReq UserRegistrationRequest

	checkRequired(w, r, &userReq)

    bytes,_ := bcrypt.GenerateFromPassword([]byte(userReq.Password), 14)

    user := User{
        ID:       uuid.New(),
        Username: userReq.Username,
        Password: string(bytes),
        CreatedAt: time.Now(),
    }

    user.Pfpfile = user.ID.String();

    
    var existingUser User
    err = database.Where("username = ?", user.Username).First(&existingUser).Error;
    if err == nil {
        http.Error(w, "Username already exists", http.StatusConflict)
        return
    }

    err = database.Create(&user).Error;
    if err != nil {
        http.Error(w, "Failed to create user", http.StatusInternalServerError)
        return
    }

    saveProfilePicture(w, r, &user)


    token, errNum := createJWT(&user)
    if errNum == 500 {
        w.WriteHeader(errNum)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    w.WriteHeader(http.StatusCreated)
    token.WriteTo(w)
}

func saveProfilePicture(w http.ResponseWriter,r *http.Request, user *User){
	file, _, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Error retrieving file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    uploadsDir := "/uploads"

    if err := os.MkdirAll(uploadsDir, os.ModePerm); err != nil {
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
}

func checkRequired(w http.ResponseWriter, r *http.Request, userReq *UserRegistrationRequest) {
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
}