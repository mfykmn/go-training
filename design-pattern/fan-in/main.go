package main

import (
	"reflect"
	"fmt"
)

func main() {
	dst := make(chan interface{})
	src1 := make(chan interface{})
	src2 := make(chan interface{})
	var srcs = []chan interface{}{src1, src2}

	go fanin(srcs, dst)()

	go display(dst)

	src1 <- "src1: aaa"
	src1 <- "src1: bbb"
	src2 <- "src2: ccc"
	src1 <- "src1: ddd"
	src2 <- "src2: eee"


}


func fanin(srcs []chan interface{}, dst chan interface{}) func() {
	cases := []reflect.SelectCase{}
	for _, ch := range srcs {
		cases = append(cases, reflect.SelectCase{
			Dir: reflect.SelectRecv,
			Chan: reflect.ValueOf(ch),
		})
	}
	return func() {
		for {
			_, recv, _ := reflect.Select(cases)
			dst<-recv.Interface()
		}
	}
}

func display(dst <-chan interface{}) {
	for {
		select {
		case res := <-dst:
			fmt.Printf("dst: %s\n", res)
		default:
			continue
		}
	}
}