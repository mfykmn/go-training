package main

import (
	"bufio"
	"errors"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	fmt.Printf("Write: %q\n", p)
	return 0, errors.New("boom!")
}

func main() {
	w := new(Writer)
	bw := bufio.NewWriterSize(w, 3)
	bw.Write([]byte{'a'})
	bw.Write([]byte{'b'})
	bw.Write([]byte{'c'})
	bw.Write([]byte{'d'})
	err := bw.Flush()
	fmt.Println(err)
}
