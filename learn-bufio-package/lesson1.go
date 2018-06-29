package main

import (
	"fmt"
	"bufio"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Println(len(p))
	return len(p), nil
}


func main() {
	fmt.Println("Unbuffered I/O")
	w := new(Writer)
	w.Write([]byte{'a'})
	w.Write([]byte{'b'})
	w.Write([]byte{'c'})
	w.Write([]byte{'d'})
	fmt.Println("Buffered I/O")
	bw := bufio.NewWriterSize(w, 3) // 3byte超えたらFlush
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush() // 明示的に最後にFlushしてる
	if err != nil {
		panic(err)
	}
}
