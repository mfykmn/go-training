package main

import (
	"fmt"
	"time"
)

func ExampleTimer() {
	t := time.NewTimer(time.Second)
	defer t.Stop()

	var count int

	for {
		select {
		case <-t.C:
			fmt.Println("break")
			return
		default:
			count++
			fmt.Println(count)
		}
	}
}
func main() {
	ExampleTimer()
}
