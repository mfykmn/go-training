package main

import (
	"bytes"
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