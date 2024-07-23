package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
