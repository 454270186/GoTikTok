package middleware

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := c.Query("token")
		if tokenStr == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{
					"status_code": -1,
					"status_msg": "missing token",
				},
			)
		} else {
			token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
				secretKey := []byte(auth.HMACSECRET)

				if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("unexpect signing method")
				}
				
				return secretKey, nil
			})

			if err != nil || !token.Valid {
				c.AbortWithStatusJSON(http.StatusUnauthorized,
					gin.H{
						"status_code": -1,
						"status_msg":  "invalid token",
					},
				)
				
				return
			}
	
			mapClaim, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("error while accessing mapclaims")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status_code": -1,
					"status_msg":  "unexpected internal error",
				})
				return
			}
	
			idFromToken := mapClaim["id"]
			userid, err := strconv.ParseFloat(c.Query("user_id"), 64)
			if err != nil {
				c.AbortWithError(http.StatusBadRequest, err)
				return
			}
	
			if idFromToken != userid {
				log.Println("user id does not match")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status_code": -1,
					"status_msg":  "user id does not match",
				})
				return
			}
	
			c.Next()
		}
	}
}