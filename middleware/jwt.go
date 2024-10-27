package middleware

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// Set jwtKey from environment variable
var jwtKey []byte

// Claims struct
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token
func GenerateJWT(userID int) (string, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load the JWT secret from environment variables
	jwtKey = []byte(os.Getenv("JWT_SECRET"))

	// Check if JWT_SECRET is set and not empty
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET is not set or empty")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Create a new token with HS256 signing method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error generating token for user ID %d: %v", userID, err)
		return "", err
	}

	log.Printf("Generated token for user ID %d: %s", userID, signedToken)
	return signedToken, nil
}

// ValidateJWT checks if the token is valid and returns the claims
func ValidateJWT(tokenString string) (*Claims, error) {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Load the JWT secret from environment variables
	jwtKey = []byte(os.Getenv("JWT_SECRET"))

	// Check if JWT_SECRET is set and not empty
	if len(jwtKey) == 0 {
		log.Fatal("JWT_SECRET is not set or empty")
	}

	// Check if the token is blacklisted
	if IsTokenBlacklisted(tokenString) {
		log.Println("Token is blacklisted")
		return nil, fmt.Errorf("blacklisted token")
	}

	// Parse the token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method matches the expected signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})

	if err != nil {
		log.Printf("Error parsing token: %v", err)
		return nil, err
	}

	// Check if the token is valid and extract claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		log.Printf("Token is valid for user ID: %d", claims.UserID)
		return claims, nil
	}

	log.Println("Invalid token provided")
	return nil, fmt.Errorf("invalid token")
}

// InvalidateJWT adds a JWT to the blacklist
func InvalidateJWT(tokenString string) {
	if tokenString == "" {
		log.Println("No token provided for invalidation")
		return
	}

	// Blacklist the token
	BlacklistToken(tokenString)
	log.Printf("Token invalidated: %s", tokenString)
}
