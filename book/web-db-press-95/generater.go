package main

import (
	"fmt"
	"sync"
)

func GenerateSeries(max int) <-chan int {
	c := make(chan int)
	go func() {
		defer close(c)
		for i := 0; i < max; i++ {
			c <- i
		}
	}()
	return c
}
func ReadSeries(wg *sync.WaitGroup, c <-chan int) {
	defer wg.Done()
	for i := range c {
		fmt.Println("Read %d", i)
	}
}

func ExampleGenerator() {
	c := GenerateSeries(100)

	var wg sync.WaitGroup
	wg.Add(2)
	go ReadSeries(&wg, c)
	go ReadSeries(&wg, c)

	wg.Wait()
}
func main() {
	ExampleGenerator()
}
