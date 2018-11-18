package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", JwtMiddleware.Handler(http.HandlerFunc(handler)))
	http.HandleFunc("/auth", GetTokenHandler)

	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world"))
}
