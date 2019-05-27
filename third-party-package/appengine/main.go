package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/cloudscheduler/v1"
)

var (
	projectID = os.Getenv("PROJECT_ID")
	locationID = os.Getenv("LOCATION_ID")
)

func main() {
	db, err := newDB()
	if err != nil {
		panic(err)
	}



	ctx := context.Background()
	c, err := google.DefaultClient(ctx, cloudscheduler.CloudPlatformScope)
	if err != nil {
		panic(err)
	}
	cloudschedulerService, err := cloudscheduler.New(c)
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

		r.Get("/schedule", func(w http.ResponseWriter, r *http.Request) {
			parent := "projects/"+projectID+"/locations/"+locationID
			rb := &cloudscheduler.Job{
				Description: "Created by GAE",
				AppEngineHttpTarget: &cloudscheduler.AppEngineHttpTarget{
					HttpMethod: http.MethodGet,
					RelativeUri: "/admin/db",
				},
				TimeZone: "Asia/Tokyo",
				Schedule: "* * * * *",
			}

			resp, err := cloudschedulerService.Projects.Locations.Jobs.Create(parent, rb).Context(ctx).Do()
			if err != nil {
				http.Error(w, fmt.Sprintf("Failed create schdule job: %v", err), 500)
				return
			}

			buf := bytes.NewBufferString(fmt.Sprintf("%#v\n", resp))
			w.Write(buf.Bytes())
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