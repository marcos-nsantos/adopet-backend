package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = os.Getenv("SECRET_KEY")

func GenerateToken(userID uint64, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   userID,
		"user_type": userType,
		"exp":       time.Now().Add(time.Hour * 1).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}
