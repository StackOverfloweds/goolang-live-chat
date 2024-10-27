package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Set jwtKey from environment variable
var jwtKey = []byte(os.Getenv("JWT_SECRET"))

//claims struct
type Claims struct {
	UserID int `json:"user_id"`
	jwt.RegisteredClaims
}

//GenerateJWT
func GenerateJWT (userID int) (string, error) {
	expirationTime := time.Now().Add(24*time.Hour)
	Claims := &Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodES256, Claims)
	return token.SignedString(jwtKey)
}

//validate JWT
func ValidateJWT (tokenSTR string) (*Claims, error) {
	Claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenSTR, Claims, func (token *jwt.Token) (interface{}, error){
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			return nil, err
		}
		return Claims, nil
	}

