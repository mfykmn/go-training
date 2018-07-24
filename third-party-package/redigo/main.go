package main

import (
	"net/http"
	"log"

	"github.com/mafuyuk/go-training/third-party-package/redigo/redis"
)

func main() {
	client, err := redis.New("redis:6379")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	http.HandleFunc("/",handler)
	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world"))
}
