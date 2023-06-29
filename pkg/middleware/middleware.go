package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/454270186/GoTikTok/pkg/auth"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// Verify token
func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.Path == "/douyin/feed/" {
			c.Next()
			return
		}

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
	
			_, ok := token.Claims.(jwt.MapClaims)
			if !ok {
				log.Println("error while accessing mapclaims")
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"status_code": -1,
					"status_msg":  "unexpected internal error",
				})
				return
			}
			
			c.Next()
		}
	}
}