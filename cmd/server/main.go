package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/sudovishal/go-url-shortener/internal/db"
	"github.com/sudovishal/go-url-shortener/internal/routes"
)

func main() {
	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	mux := routes.SetupRoutes()
	log.Println("Server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
