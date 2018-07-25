package main

import (
	"fmt"
)

func main() {
	reader, exit := read()
	stringSlice := []string{"a", "b", "c"}

	for _, v := range stringSlice {
		reader <- v
	}

	exit <- struct{}{}
}


func read() (reader chan string, exit chan struct{}) {
	// 値の受け渡し用のチャネル
	reader = make(chan string, 10)
	// 終了フラグを渡すチャネル
	exit = make(chan struct{})
	go func() {
		for {
			select {
			case data := <-reader:
				fmt.Println(data)
			case <-exit:
				fmt.Println("exit")
				break
			}
		}
	}()
	return reader, exit
}
