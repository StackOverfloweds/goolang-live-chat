package Auth

import (
	"encoding/json"
	"go-live-chat/middleware"
	"go-live-chat/model"
	"log"
	"net/http"
	"strings"
)

func LoginController(w http.ResponseWriter, r *http.Request) {
	// Check if the user is already logged in by verifying the Bearer token
	tokenSTR := r.Header.Get("Authorization")
	if tokenSTR != "" {
		// Check for the bearer
		if !strings.HasPrefix(tokenSTR, "Bearer ") {
			http.Error(w, "Invalid Authorization format", http.StatusUnauthorized)
			return
		}

		// Extract token 
		tokenSTR = strings.TrimPrefix(tokenSTR, "Bearer ")

		// Validate the token
		claims, err := middleware.ValidateJWT(tokenSTR)
		if err == nil && claims != nil {
			log.Printf("User is already logged in. User ID: %d", claims.UserID)
			http.Error(w, "User is already logged in", http.StatusForbidden)
			return
		} else {
			log.Printf("Token validation failed: %v", err)
		}
	}

	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	// Decode the request body to get email and password
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}
	// Authenticate the user
	user, err := model.Authenticate(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Invalid Credential", http.StatusUnauthorized)
		return 
	}
	// Generate JWT token for the authenticated user
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	// Store the token in the middleware
	middleware.StoreToken(token, user.ID)

	// Send both tokens to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token, "user": user.Username})
}
