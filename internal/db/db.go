package db

import (
	"database/sql"
	"fmt"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func Connect(url, token string) (*sql.DB, error) {

	dsn := fmt.Sprintf("%s?authToken=%s", url, token)
	db, err := sql.Open("libsql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening Database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error connecting to database: %v", err)
	}

	fmt.Println("Connected to database")
	return db, nil
}
