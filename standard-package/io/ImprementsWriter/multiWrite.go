package main

import (
	"net/http"
	"encoding/json"
	"compress/gzip"
	"io"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello":"World",
	}

	bytes, err := json.Marshal(source)
	if err != nil {
		return
	}

	writer := gzip.NewWriter()
	writer.Flush()

	w.Write()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
