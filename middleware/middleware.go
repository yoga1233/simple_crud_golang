package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mendapatkan token dari header Authorization
		tokenHeader := c.GetHeader("Authorization")

		// Memeriksa apakah token ada
		if tokenHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Missing token"})
			c.Abort()
			return
		}

		// Memeriksa apakah token memiliki format yang benar
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token format"})
			c.Abort()
			return
		}

		// Mengambil token tanpa 'Bearer '
		tokenValue := splitted[1]

		fmt.Println("ini adalah value token", tokenValue)

		// Memeriksa dan mendekode token
		token, err := jwt.Parse(tokenValue, func(token *jwt.Token) (interface{}, error) {
			// Ganti dengan secret key yang sesuai dengan yang digunakan untuk menghasilkan token
			return []byte(os.Getenv("JWT_KEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: Invalid token"})
			c.Abort()
			return
		}

		c.Next()
	}
}
