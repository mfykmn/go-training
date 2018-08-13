package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}

	compresser := gzip.NewWriter(w)
	mulitiWriter := io.MultiWriter(compresser, os.Stdout)

	encoder := json.NewEncoder(mulitiWriter)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)

	compresser.Flush()

}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
