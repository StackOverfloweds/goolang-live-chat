package main

import (
	"go-live-chat/Helpers"
	"go-live-chat/config"
	"go-live-chat/database"
	"go-live-chat/routes"
	"log"
	"net/http"
)

func main() {

	// Ensure JWT_SECRET is present in the .env file
	Helpers.EnsureJWTSecret()

	config.Connect()
	defer config.DB.Close()

	database.CreateTables()
	log.Println("Database setup completed")

	routes.SetupRoutes()
	log.Println("server started on 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Server failed to start : ",err)
	}
}