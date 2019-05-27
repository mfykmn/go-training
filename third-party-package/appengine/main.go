package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

func main() {
	db, err := newDB()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
  scheduler, err := newScheduler(ctx)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/", helloHandler)
	r.Get("/db", func(w http.ResponseWriter, r *http.Request) {
		res, err := db.Show()
		if err != nil {
			http.Error(w, fmt.Sprintf("Failed show db: %v", err), 500)
			return
		}
		w.Write(res)
	})

	r.Get("/schedule", func(w http.ResponseWriter, r *http.Request) {
		if err := scheduler.Reserve(ctx); err != nil {
			http.Error(w, fmt.Sprintf("Failed create schdule job: %v", err), 500)
			return
		}
		fmt.Fprintln(w, "Scheduled")
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