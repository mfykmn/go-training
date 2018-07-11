package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// curl http://localhost:8080/
func main() {
	r := chi.NewRouter()
	// r.Use(middleware.RequestID) TODO
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		rqqID := middleware.GetReqID(ctx)

		w.Write([]byte(`{"request_id":"`+rqqID+`"}`))
	})
	http.ListenAndServe(":8080", r)
}
