package config

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Load environment variables
var JwtSecret = []byte(GetEnv("JWT_KEY", "JWT_KEY"))

// GenerateJWT generates a new JWT token
func GenerateJWT(userID uint, email string) (string, error) {
	fmt.Println(JwtSecret)

	// Define expiration time (e.g., 24 hours)
	expirationTime := time.Now().Add(24 * time.Hour)

	// Create JWT claims
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     expirationTime.Unix(),
	}

	// Create a new token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	return token.SignedString(JwtSecret)
}

// VerifyJWT verifies and parses the JWT token
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return JwtSecret, nil
	})
}
