package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func ExampleWaitGroup() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			n := rand.Intn(5)
			fmt.Println("Going to sleep for", n, "seconds")
			time.Sleep(time.Duration(n) * time.Second)
		}()
	}

	wg.Wait() // 終了を待つ
	fmt.Println("All goroutines finishied")
}

func main() {
	ExampleWaitGroup()
}
