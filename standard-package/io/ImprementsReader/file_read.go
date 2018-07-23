package main

import (
	"os"
	"io"
)

// go run file_read.go
func main() {
	file, err := os.Open("file_read.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
