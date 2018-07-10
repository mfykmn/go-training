package main

import (
	"sync"
	"fmt"
	"context"
)

var wg sync.WaitGroup


func main() {
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURL(ctx, queue)
	}


	queue <- "http://www.example.com"
	queue <- "http://www.example.net"
	queue <- "http://www.example.net/foo"
	queue <- "http://www.example.net/bar"

	cancel()
	wg.Wait()
}

func fetchURL(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exit")
			wg.Done()
			return
		case url := <- queue:
			fmt.Println("fetching", url)
		}
	}
}
