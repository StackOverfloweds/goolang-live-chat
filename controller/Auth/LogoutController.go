package Auth

import (
	"go-live-chat/middleware"
	"log"
	"net/http"
	"strings"
)

func LogoutController (w http.ResponseWriter, r * http.Request){
	//get the Authorization header
	tokenSTR := r.Header.Get("Authorization")
	if tokenSTR == "" {
		http.Error(w, "Authorization header is missing", http.StatusUnauthorized)
		return
	}

	// Check for the Bearer prefix
	if !strings.HasPrefix(tokenSTR, "Bearer ") {
		http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
		return
	}

	// Extract the token
	tokenSTR = strings.TrimPrefix(tokenSTR, "Bearer ")

	// Blacklist the token
	middleware.BlacklistToken(tokenSTR)

	// log the logout action
	log.Printf("User logged out. Token has been blacklisted: %s", tokenSTR)

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfully logged out"))

}