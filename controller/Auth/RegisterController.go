package Auth

import (
	"encoding/json"
	"go-live-chat/model"
	"net/http"
)

//create login controller handler
func RegisterController (w http.ResponseWriter, r *http.Request) {
	var user model.User

	if err :=json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid input", http.StatusBadRequest)
		return
	}

	if err := user.Register(); err != nil {
		http.Error(w, "Error Registered User", http.StatusInternalServerError)
		return
	}

	w.Write([]byte("user Registered Successfully"))
}