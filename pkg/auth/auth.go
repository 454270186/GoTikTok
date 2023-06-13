package auth

import (
	"errors"
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

/*
	Password Crypting
*/
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

/*
	Token Parser
*/
func GetUIDFromToken(tokenStr string) (uint, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		secretKey := []byte(HMACSECRET)

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpect signing method")
		}

		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	mapClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("error while accessing mapclaims")
	}

	return mapClaims["id"].(uint), nil
}