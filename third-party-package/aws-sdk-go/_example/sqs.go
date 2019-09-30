package main

import (
	"go-aws-mock"
	"log"
)

func main() {
	// 動作確認のため同期で動かしているが
	// goroutineで非同期に動かして、channelでデータをやり取りするべし

	sess := mock.InitAWSSession()
	sqs := mock.NewSimpleQueueService(sess)

	if err := sqs.SendMessage("{key:val}"); err != nil {
		log.Println(err)
	}

	func() {
		res, err := sqs.ReceiveMessage()
		log.Println("*** ReceiveMessage 1 ***")
		log.Println(res)
		log.Println(err)
	}()

	func() {
		res, err := sqs.ReceiveMessage()
		log.Println("*** ReceiveMessage 2 ***")
		log.Println(res)
		log.Println(err)
	}()
}
