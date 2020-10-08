package main

import (
	"fmt"
	"sync"
	"time"
)

func ExampleCond() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer mu.Unlock()

			fmt.Printf("waiting %d\n", i)
			mu.Lock()
			c.Wait()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	for i := 0; i < 10; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Printf("signaling!\n")
		// 通知
		c.Signal()
	}

	wg.Wait()
}

func ExampleBroadcastCond() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	c := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			defer mu.Unlock()

			fmt.Printf("waiting %d\n", i)
			mu.Lock()
			c.Wait()
			fmt.Printf("go %d\n", i)
		}(i)
	}

	// 通知
	time.AfterFunc(time.Second, c.Broadcast)

	wg.Wait()
}

func main() {
	fmt.Println("***********ExampleCond***********")
	ExampleCond()

	fmt.Println("*******ExampleBroadcastCond*******")
	ExampleBroadcastCond()
}
