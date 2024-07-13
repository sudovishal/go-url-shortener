package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/sudovishal/go-url-shortener/internal/db"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get environment variables
	dbURL := os.Getenv("DB_URL")
	dbToken := os.Getenv("DB_TOKEN")

	// Connect to TursoDB
	db, err := db.Connect(dbURL, dbToken)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Database Connection Established")

	log.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
