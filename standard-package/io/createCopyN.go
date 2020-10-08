package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// Goならわかるシステムプログラミングの回答 Q3.5
func main() {
	reader1 := strings.NewReader("mfykmn")
	reader2 := strings.NewReader("mfykmn")

	// もとからあるCopyN
	written1, err := io.CopyN(os.Stdout, reader1, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(written1)

	// 自前実装のCopyN
	written2, err := CopyN(os.Stdout, reader2, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println(written2)
}

func CopyN(dst io.Writer, src io.Reader, n int64) (written int64, err error) {
	reader := io.LimitedReader{src, n}
	return io.Copy(dst, &reader)
}
