package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mafuyuk/go-training/third-party-package/redigo/redis"
)

func main() {
	client := redis.New("redis:6379")
	sconn := client.Pool.Get()
	defer sconn.Close()

	msgRepo := redis.NewMsg(sconn)
	go msgRepo.Receive()

	sh := sub{
		msgRepo: msgRepo,
	}

	ph := pub{
		client,
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/sub", sh.subHandler)
	http.HandleFunc("/pub", ph.pubHandler)
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

type pub struct {
	*redis.Client
}

func (p *pub) pubHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("call pubhandler")
	conn := p.Pool.Get()
	defer conn.Close()

	if err := conn.Send("PUBLISH", channel, "hello pubsub"); err != nil {
		fmt.Println(err)
		rw.Write([]byte("failed pubhandler"))
		return
	}
	rw.Write([]byte("send message"))
}
