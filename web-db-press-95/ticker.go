package main

import (
	"fmt"
	"time"
)

func ExampleTicker() {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	for i := 0; i < 10; i++ {
		select {
		case <-t.C:
			fmt.Printf("%d\n", i)
		}
	}
}

func ExampleSelectTicker() {
	t := time.NewTicker(time.Second)
	defer t.Stop()

	in := make(chan struct{})

	var c chan struct{}

	for {
		select {
		case <-t.C:
			c = in // c != nilだと次のcaseが読み込める
		case <-c:
			c = nil // c == nilだと"<-c"はブロックし続ける
		}
	}
}

func main() {
	fmt.Println("********ExampleTicker********")
	ExampleTicker()

	fmt.Println("*****ExampleSelectTicker*****")
	//ExampleSelectTicker()

}
