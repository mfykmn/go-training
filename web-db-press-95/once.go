package main

import (
	"fmt"
	"sync"
)

var Gloval int
var initGlobal sync.Once

func ExampleOnce() {
	cb := func(wg *sync.WaitGroup) {
		defer wg.Done()
		initGlobal.Do(func() {
			// Globalの初期化
			Gloval = 100
			fmt.Println("init Global", Gloval)
		})
		Gloval++
		fmt.Println("increment Global", Gloval)
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go cb(&wg)
	}
	wg.Wait()
}
func main() {
	ExampleOnce()
}
