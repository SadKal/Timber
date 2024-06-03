package db

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
)

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

func GetUserByUsername(username string, database *gorm.DB) (*User, error) {
	var user User

	if err := database.Preload("Chats").First(&user, "username = ?", username).Error; err != nil {
		return nil, err
	}

	return &user, nil
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

