package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var port = os.Getenv("PORT")

func main() {
	// Init
	db, err := newDB()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
  scheduler, err := newScheduler(ctx)
	if err != nil {
		panic(err)
	}

  // Routing
	r := chi.NewRouter()
	r.Get("/", helloHandler)
	r.Get("/db", dbHandlerFunc(db))
	r.Get("/db/schedule", scheduleHandlerFunc(ctx, scheduler))

	// Serve
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}