package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

// Connect initializes the database connection
func Connect () {

	// Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", username, password, dbName)
	db, err := sql.Open("mysql",dsn)
	if err != nil {
		log.Fatal("error connection to database : ",err)
	}

	//verify the connection
	if err :=db.Ping(); err != nil {
		log.Fatal("could not ping database : ", err)
	}

	DB = db
	log.Printf("Database Connected")
}
