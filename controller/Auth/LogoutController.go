package auth

import "net/http"

func LogoutController (w http.ResponseWriter, r * http.Request){
	w.Write([]byte("User Logged Out Succesfully"))
}