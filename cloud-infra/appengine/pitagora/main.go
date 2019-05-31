package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

var (
	port = os.Getenv("PORT")
	projectID = os.Getenv("PROJECT_ID")
	locationID = os.Getenv("LOCATION_ID")
)

func main() {
	// Init
	db, err := newDB()
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
  scheduler, err := newScheduler(ctx, projectID, locationID)
	if err != nil {
		panic(err)
	}
	storage, err := newStorage(ctx, projectID)
	if err != nil {
		panic(err)
	}

  // Routing
	r := chi.NewRouter()
	r.Get("/", helloHandler)
	r.Get("/db", dbHandlerFunc(db))
	r.Get("/schedule", scheduleHandlerFunc(ctx, scheduler))
	r.Get("/storage/download", storageDownloadHandlerFunc(ctx, storage))
	r.Get("/storage/upload", storageUploadHandlerFunc(ctx, storage))

	// Serve
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}