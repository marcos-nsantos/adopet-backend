package auth

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token not provided"})
			return
		}

		authHeaderParts := strings.Split(authHeader, " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		tokenString := authHeaderParts[1]
		valid, err := validateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if !valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token is not valid"})
			return
		}

		userID, err := getUserIDFromToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		userType, err := getUserTypeFromToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Set("user_id", userID)
		c.Set("user_type", userType)
		c.Next()
	}
}

func validateToken(tokenString string) (bool, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return false, err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return true, nil
	}

	return false, nil
}

func parseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return []byte(secretKey), nil
	})
}

func getUserIDFromToken(tokenString string) (uint64, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, err
	}

	userIDFloat := claims["user_id"].(float64)
	userID := uint64(userIDFloat)

	return userID, nil
}

func getUserTypeFromToken(tokenString string) (string, error) {
	token, err := parseToken(tokenString)
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", err
	}

	userType := claims["user_type"].(string)

	return userType, nil
}
