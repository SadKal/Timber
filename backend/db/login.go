package db

import (
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func LoginUser(w http.ResponseWriter, r *http.Request, database *gorm.DB) {
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
    err = database.Where("BINARY username = ?", user.Username).First(&dbUser).Error
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