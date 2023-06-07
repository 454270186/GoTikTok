package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	HMACSECRET = "xiaofei"
	TOKEN_DURATION = 24 * time.Hour
)
func NewTokenByUserID(userId uint) (string, string) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": userId,
		"exp": time.Now().Add(TOKEN_DURATION).Unix(),
	})

	tokenSignedStr, err := token.SignedString([]byte(HMACSECRET))
	if err != nil {
		log.Println(err)
		return "", ""
	}

	return tokenSignedStr, ""
}