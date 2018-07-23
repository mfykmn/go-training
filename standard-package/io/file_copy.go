package main

import (
	"os"
	"io"
)

// go run file_copy.go
func main() {
	src, err := os.Open("new.txt")
	if err != nil {
		panic(err)
	}
	defer src.Close()

	dst, err := os.OpenFile("old.txt", os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer dst.Close()

	io.Copy(dst, src)
}
