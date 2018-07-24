package main

import (
	"flag"
	"net/http"
	"log"

	"github.com/mafuyuk/go-training/third-party-package/redigo/redis"
)

type option struct {
	host string
	port string
}

func main() {
	var opt = &option{}
	flag.StringVar(&opt.host, "h", "", "Server host")
	flag.StringVar(&opt.port, "p", "", "Server port")
	flag.Parse()

	client, err := redis.New("redis:6379")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	http.HandleFunc("/",handler)
	log.Println("Start server")
	http.ListenAndServe(opt.host+opt.port, nil)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world"))
}
