package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"cloud.google.com/go/scheduler/apiv1"
	schedulerpb "google.golang.org/genproto/googleapis/cloud/scheduler/v1"
)

func main() {
	db, err := newDB()
	if err != nil {
		panic(err)
	}

	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()
	c, err := scheduler.NewCloudSchedulerClient(ctx)
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()
	r.Get("/", helloHandler)
	r.Route("/admin", func(r chi.Router) {
		r.Get("/db", func(w http.ResponseWriter, r *http.Request) {
			rows, err := db.Query("SHOW DATABASES")
			if err != nil {
				http.Error(w, fmt.Sprintf("Could not query db: %v", err), 500)
				return
			}
			defer rows.Close()

			buf := bytes.NewBufferString("Databases:\n")
			for rows.Next() {
				var dbName string
				if err := rows.Scan(&dbName); err != nil {
					http.Error(w, fmt.Sprintf("Could not scan result: %v", err), 500)
					return
				}
				fmt.Fprintf(buf, "- %s\n", dbName)
			}
			w.Write(buf.Bytes())
		})

		r.Post("/schedule", func(w http.ResponseWriter, r *http.Request) {
			req := &schedulerpb.CreateJobRequest {
				Parent: "projects/"+projectID+"/locations/asia-northeast1",
				Job: &schedulerpb.Job {
					Name: "test",
					Description: "test",
					Schedule: "* * * * *",
					TimeZone: "timeZone36848094",
				},
			}
			resp, err := c.CreateJob(ctx, req)
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed CreateJob: %v", err), 500)
				return
			}

			_ = resp
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