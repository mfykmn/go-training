package main

import (
	"log"

	"github.com/sfreiberg/gotwilio"
)

// see https://github.com/sfreiberg/gotwilio
const (
	accountSid = ""
	authToken  = ""

	from = "+819000000000"                          // 受信するtwilio側の電話番号
	to   = "+819000000000"                          // 送信する電話番号
	url  = "https://handler.twilio.com/twiml/dummy" // twiml URL
)

func main() {
	twilio := gotwilio.NewTwilioClient(accountSid, authToken)
	callbackParams := gotwilio.NewCallbackParameters(url)
	_, exc, err := twilio.CallWithUrlCallbacks(from, to, callbackParams)

	if err != nil {
		log.Fatal(err)
	}

	if exc != nil {
		log.Fatal(exc)
	}
}
