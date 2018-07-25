package main

import "fmt"

func main() {
	// バッファなし
	tasks1 := make(chan string)
	fmt.Printf("%#v\n", tasks1)
	// バッファ付き
	tasks2 := make(chan string, 10)
	fmt.Printf("%#v\n", tasks2)

	go func() {
		// データを送信
		tasks1 <- "cmake .."
		tasks1 <- "cmake . --build Debug"
		tasks1 <- "hello"
	}()


	// データを受け取り
	task1 := <- tasks1
	fmt.Println(task1)

	// データを受け取り＆クローズ判定
	if task2, ok := <- tasks1; ok {
		fmt.Println(task2)
	}

	// データ読み捨て
	<- tasks1
}
