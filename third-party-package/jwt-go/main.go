package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world"))
}