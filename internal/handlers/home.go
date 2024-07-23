package handlers

import (
	// "html/template"
	"net/http"

	"github.com/sudovishal/go-url-shortener/templates"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	templates.Home().Render(r.Context(), w)
}
