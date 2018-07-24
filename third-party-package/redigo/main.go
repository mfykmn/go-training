package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mafuyuk/go-training/third-party-package/redigo/redis"
)

func main() {
	var err error
	client, err := redis.New("redis:6379")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	msgRepo := redis.NewMsg(client)
	defer msgRepo.Close() // TODO タイミングが違う？
	go msgRepo.Receive()

	sh := sub{
		msgRepo: msgRepo,
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/sub", sh.subHandler)
	http.HandleFunc("/pub", sh.pubHandler)
	log.Println("Start server")
	http.ListenAndServe(":8080", nil)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Hello world"))
}

type sub struct {
	msgRepo *redis.MessageRepository
}

var channel = "dummy"

func (s *sub) subHandler(rw http.ResponseWriter, r *http.Request) {
	if err := s.msgRepo.Sub(channel); err != nil {
		fmt.Println(err)
		rw.Write([]byte("failed subhandler"))
		return
	}
	rw.Write([]byte("exec subscription"))
}

func (s *sub) pubHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("call pubhandler")
	if err := s.msgRepo.Pub(channel, "hello pubsub"); err != nil {
		fmt.Println(err)
		rw.Write([]byte("failed pubhandler"))
		return
	}
	rw.Write([]byte("send message"))
}
