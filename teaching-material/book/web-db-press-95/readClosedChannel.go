package main

import "fmt"

func ExampleReadClosedChannel() {
	ch := make(chan int, 10)
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

	for i := 0; i < 10; i++ {
		v := <-ch // 1から10まで読み込まれる
		fmt.Println("Read", v)
	}

	// Pattern 1
	func() {
		v := <-ch              // i = 0, intのゼロ値は0
		fmt.Println("Read", v) // ゼロ値だったらchannelは空
	}()

	// Pattern 2
	func() {
		v, ok := <-ch // 第２引数をとった場合は第２引数がfalseだったらchannelは空
		if !ok {
			fmt.Println("Read", v)
		}
	}()
}

func main() {
	ExampleReadClosedChannel()
}
