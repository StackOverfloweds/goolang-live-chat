package auth

import (
	"encoding/json"
	"go-live-chat/middleware"
	"go-live-chat/model"
	"net/http"
)

func LoginController(w http.ResponseWriter, r * http.Request) {
	var creds struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	if err :=json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "invalid input",http.StatusBadRequest)
		return
	}

	user, err:=model.Authenticate(creds.Email,creds.Password)
	if err!=nil {
		http.Error(w, "Invalid Credential",http.StatusUnauthorized)
		return 
	}

	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		http.Error(w, "could not generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "aplication/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}