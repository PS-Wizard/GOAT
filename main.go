package main

import (
	"log"
	"net/http"

	"github.com/PS-Wizard/GOAT/views/pages"
	"github.com/a-h/templ"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
	http.Handle("/", templ.Handler(pages.Home()))
	http.Handle("/about", templ.Handler(pages.About()))

	log.Println("Listening on :8080")
	http.ListenAndServe(":8080", nil)
}
