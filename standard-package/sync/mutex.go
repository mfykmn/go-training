package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			counter++
			fmt.Print("*")
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("\n%d\n", counter)
}
