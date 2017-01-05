package main

import (
	"fmt"
	"sync"
)

func printString(sem chan struct{}, str string) {
	fmt.Println("start", str)
	sem <- struct{}{}
	defer func() {
		fmt.Println("lerease semaphore", str)
		<-sem
	}() // 解放
}

func ExampleSemaphore() {
	sem := make(chan struct{}, 3)
	stringSlice := []string{"a", "b", "c", "d", "e", "f", "g"}
	var wg sync.WaitGroup
	for _, v := range stringSlice {
		str := v // ポインタを渡しているのでコピーしないと同じ値になってしまう
		wg.Add(1)
		go func() {
			defer wg.Done()
			printString(sem, str)
		}()
	}
	wg.Wait()
}
func main() {
	ExampleSemaphore()
}
