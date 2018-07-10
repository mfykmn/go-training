package main

import (
	"sync"
	"fmt"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(queue)
	}


	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.net/foo"
	queue <- "http://www.example.net/bar"

	close(queue)
	wg.Wait()
}

func fetchURL(queue chan string) {
	for {
		url, more := <-queue // closeされるとmoreがfalseになる
		if more {
			fmt.Println("fetching", url)
		} else {
			fmt.Println("worker exit")
			wg.Done()
			return
		}
	}
}
