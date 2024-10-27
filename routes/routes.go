package routes

import (
	"encoding/json"
	"go-live-chat/controller/Auth"
	"net/http"
)

// setup routes
func SetupRoutes () {
	//create new serveMux
	mux :=http.NewServeMux()

	// Define a health check route at root ("/") to return "Server ready" message
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Server Ready"})
	})

	//define the routes
	mux.HandleFunc("/Auth/register", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Auth.RegisterController(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/Auth/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Auth.LoginController(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/Auth/logout", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			Auth.LogoutController(w,r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	//start the HTTP server with mux handler
	http.Handle("/",mux)

}