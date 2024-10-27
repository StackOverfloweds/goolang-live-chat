package Helpers

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"strings"
)

// GenerateRandomSecret generates a random secret key of the given length
func GenerateRandomSecret(length int) (string, error) {
	secret := make([]byte, length)
	_, err := rand.Read(secret)
	if err != nil {
		return "", fmt.Errorf("failed to generate random secret: %w", err)
	}
	return base64.StdEncoding.EncodeToString(secret), nil
}

// EnsureJWTSecret checks for JWT_SECRET in the .env file and generates it if not present
func EnsureJWTSecret() {
	envFile := ".env"
	const secretKey = "JWT_SECRET"

	// Check if the JWT_SECRET already exists in the .env file
	if _, err := os.Stat(envFile); err == nil {
		fileContent, err := os.ReadFile(envFile)
		if err != nil {
			log.Fatalf("Error reading .env file: %v", err)
		}

		if strings.Contains(string(fileContent), secretKey) {
			log.Println("JWT_SECRET already exists in .env file.")
			return
		}
	}

	// Generate a new JWT secret
	secret, err := GenerateRandomSecret(32)
	if err != nil {
		log.Fatal("Error generating JWT secret:", err)
	}

	// Save the secret to the .env file
	file, err := os.OpenFile(envFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening .env file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(fmt.Sprintf("%s=%s\n", secretKey, secret)); err != nil {
		log.Fatalf("Error writing to .env file: %v", err)
	}

	log.Println("JWT_SECRET generated and saved to .env file.")
}
