package routes

import (
	"net/http"

	"github.com/sudovishal/go-url-shortener/internal/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", handlers.HomeHandler)
	return mux
}
