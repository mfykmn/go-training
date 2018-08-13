package main

import (
	"bytes"
	"fmt"
	"io"
	"time"
)

func main() {
	r, w := io.Pipe()

	// 同期的にするとデッドロックが起こることがあるので goroutineで並列処理にするか、bufioでバッファーさせないといけない
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutine")
		fmt.Fprint(w, "some text to be read\n")
		w.Close()
	}()
	buf := new(bytes.Buffer)
	buf.ReadFrom(r) // writeされるまでまちになる
	fmt.Println("main")
	fmt.Print(buf.String())
}
