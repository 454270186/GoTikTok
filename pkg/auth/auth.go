package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
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

func GetHashedPwd(pwd string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return hash, nil
}

func ComparePwd(hashedPwd string, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd))
	if err != nil {
		return false
	} else {
		return true
	}
}