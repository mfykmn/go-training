package main

import (
	"fmt"
	"sync"
	"time"
)

func ExampleRWMutex() {
	var mu sync.RWMutex
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		id := i
		go func() {
			defer wg.Done()
			for i := 0; i < 5; i++ {
				// 共有ロックを得る
				mu.RLock()
				fmt.Printf("Reader %d: Acquired lock\n", id)
				time.Sleep(time.Second)
				mu.RUnlock()
			}
		}()
	}

	time.AfterFunc(3*time.Second, func() {
		fmt.Println("Writer: Acquired lock")
		// 排他ロックを得る
		mu.Lock()
	})

	time.AfterFunc(6*time.Second, func() {
		fmt.Println("Writer: Releasing lock")
		// 排他ロックを解放する
		mu.Unlock()
	})

	wg.Wait()
}

func main() {
	ExampleRWMutex()
}
