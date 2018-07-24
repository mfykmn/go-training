package main

import (
	"os"
	"crypto/rand"
	"io"
)

// GoならわかるシステムプログラミングQ3.2の自分の回答
// 実際の答えは randRead2.go のほう
func main() {
	file, err := os.Create("rand.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	limitReader := io.LimitedReader{
		R: rand.Reader,
		N: 1024,
	}
	io.Copy(file, &limitReader)
}
