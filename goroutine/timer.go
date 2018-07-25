package main

import (
	"time"
	"fmt"
)

func main() {
	fmt.Println("timer start")
	duration := 3 * time.Second
	timer := time.After(duration)
	<-timer
	fmt.Println(duration)
}
