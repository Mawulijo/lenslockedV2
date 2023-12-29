package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to version 2 of the great lenslocked</h1>")
}
func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Contact us at <a href=\"mailto:support@lenslocked.com\">support@lenslocked.com</a></h1>")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting Server on port 3000...")
	http.ListenAndServe(":3000", r)
}
