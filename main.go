package main

import (
	"go-live-chat/config"
	"go-live-chat/database"
	"log"
)

func main() {

	config.Connect()
	defer config.DB.Close()

	database.CreateTables()
	log.Println("Database setup completed")

}