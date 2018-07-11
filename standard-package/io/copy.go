package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("--Copy--") // すべてコピー
	reader0 := strings.NewReader("hello world!")
	size0, _ := io.Copy(os.Stdout, reader0)
	fmt.Println(size0)

	fmt.Println("--CopyN--") // 指定したサイズだけコピー
	reader1 := strings.NewReader("hello world!")
	size1, _ := io.CopyN(os.Stdout, reader1, 3)
	fmt.Println(size1)

	fmt.Println("--CopyBuffer--")
	reader2 := strings.NewReader("hello world!")
	buffer := make([]byte, 4)
	io.CopyBuffer(os.Stdout, reader2, buffer)
	fmt.Println(string(buffer))
}
