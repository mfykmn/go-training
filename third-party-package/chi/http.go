package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

// curl http://localhost:8080/
func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})
	http.ListenAndServe(":8080", r)
}
