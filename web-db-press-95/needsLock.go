package main

import (
	"fmt"
	"sync"
	"time"
)

type Object struct {
	mu sync.Mutex
	wg sync.WaitGroup
}

func (o *Object) NeedsLock() {
	o.mu.Lock()
	defer o.mu.Unlock()

	fmt.Println("A")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("B")
	time.Sleep(100 * time.Millisecond)
	fmt.Println("C")

	o.wg.Done()

}

func main() {
	obj := new(Object)

	for i := 0; i < 3; i++ {
		obj.wg.Add(1)
		go obj.NeedsLock()
	}

	obj.wg.Wait()
}
