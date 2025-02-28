package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("your-secret-key")

// GenerateJWT creates a JWT token for authentication
func GenerateJWT(userID int, email string, clientID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"email":     email,
		"client_id": clientID, // ✅ Now includes client information
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(jwtSecret)
}
