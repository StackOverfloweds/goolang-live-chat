package middleware

import (
	"sync"
)

// TokenStatus defines the status of a token
type TokenStatus int

const (
	Valid TokenStatus = iota
	Blacklisted
)

// TokenInfo holds the user ID and status of the token
type TokenInfo struct {
	UserID int
	Status  TokenStatus
}

// Global map to store token statuses
var (
	tokenStore = make(map[string]TokenInfo) // Store for tokens
	mu         sync.Mutex                   // Mutex for concurrent access
)

// BlacklistToken marks a token as blacklisted
func BlacklistToken(token string) {
	mu.Lock()
	defer mu.Unlock()
	if info, exists := tokenStore[token]; exists {
		info.Status = Blacklisted
		tokenStore[token] = info
	}
}

// IsTokenBlacklisted checks if a token is blacklisted
func IsTokenBlacklisted(token string) bool {
	mu.Lock()
	defer mu.Unlock()
	if info, exists := tokenStore[token]; exists {
		return info.Status == Blacklisted
	}
	return false
}

// StoreToken stores a token associated with a user ID and status
func StoreToken(token string, userID int) {
	mu.Lock()
	defer mu.Unlock()
	tokenStore[token] = TokenInfo{
		UserID: userID,
		Status: Valid,
	}
}

// IsTokenValid checks if a token is valid
func IsTokenValid(token string) (int, bool) {
	mu.Lock()
	defer mu.Unlock()
	if info, exists := tokenStore[token]; exists && info.Status == Valid {
		return info.UserID, true
	}
	return 0, false
}