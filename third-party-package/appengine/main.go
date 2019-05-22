package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {

		r.Get("/a", helloHandler)

		r.Route("/admin", func(r chi.Router) {
			r.Get("/", adminHandler)
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello")
}

func adminHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Admin")
}