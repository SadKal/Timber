package db

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))
var expirationTime = time.Now().Add(8 * time.Hour)


type Claims struct {
	Username string `json:"username"`
	UUID uuid.UUID `json:"uuid"`
	jwt.RegisteredClaims
}

func createJWT(user *User) (bytes.Buffer, int){

	claims := &Claims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return bytes.Buffer{}, http.StatusInternalServerError
	}

	var responseBuffer bytes.Buffer

	json.NewEncoder(&responseBuffer).Encode(map[string]interface{}{
        "message":      "Login successful",
        "token":        tokenString,
        "expiresIn":    int(time.Until(expirationTime).Seconds()),
    })

    return responseBuffer, 0;
}

func AuthToken(r *http.Request, db *gorm.DB) (Claims, int) {
	token := r.URL.Query().Get("jwt")
	claims := &Claims{}

	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return secretKey, nil
	})

	user, _ := GetUserByUsername(claims.Username, db)

	claims.UUID = user.ID

	log.Println("CLAIMS", claims.Username)
	if err != nil {
		return Claims{} , http.StatusUnauthorized
	}

    return *claims, 0;
}