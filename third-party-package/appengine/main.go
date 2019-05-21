package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Route("/v1", func(r chi.Router) {

		r.Route("/a", func(r chi.Router) {
			r.Get("/", aHandler)
		})

		r.Route("/b", func(r chi.Router) {
			r.Get("/", bHandler)
		})
	})
	http.ListenAndServe(":8080", r)
}

func aHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, A")
}

func bHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, B")
}