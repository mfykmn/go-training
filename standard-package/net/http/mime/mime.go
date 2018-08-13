package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// http://localhost:8080/csv
func csvhandler(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs("demo.csv")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment;filename=\"download.csv\"")

	buf := new(bytes.Buffer)
	io.Copy(buf, file)
	w.Write(buf.Bytes())
}

func main() {
	var httpServer http.Server
	http.HandleFunc("/csv", csvhandler)
	httpServer.Addr = ":8080"
	log.Panicln(httpServer.ListenAndServe())
}
