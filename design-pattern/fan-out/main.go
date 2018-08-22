package main

import (
	"reflect"
	"fmt"
)

// lean web db press 105

func main() {
	src := make(chan interface{})
	dst1 := make(chan interface{})
	dst2 := make(chan interface{})

	var channels = []chan interface{}{dst1, dst2}

	go fanout(src, channels)

	src <- "hello"

	fmt.Printf("dst1: %s\n", <-dst1)
	fmt.Printf("dst2: %s\n", <-dst2)
}

func fanout(src <-chan interface{}, dstChs []chan interface{}) {
	for {
		data := <- src
		cases := make([]reflect.SelectCase, len(dstChs))
		for i, ch := range dstChs {
			cases[i] = reflect.SelectCase{
				Chan: reflect.ValueOf(ch),
				Dir: reflect.SelectSend,
				Send: reflect.ValueOf(data),
			}
		}
		for len(cases) > 0 {
			chosen, _, _ := reflect.Select(cases)
			cases = append(cases[:chosen], cases[chosen+1:]...)
		}
	}
}