package main

import (
	"net/http"
	"io"
)

func handler(w http.ResponseWriter, r * http.Request) {
	io.WriteString(w, "http.ResponsWriter sample")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
