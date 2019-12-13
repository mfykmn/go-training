package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	fmt.Print("0")
	for conditionTrue() == false {
		c.Wait()
	}
	fmt.Print("3")
	c.L.Unlock()
}

func conditionTrue() bool {
	time.Sleep(2 + time.Second)
	return true

}
