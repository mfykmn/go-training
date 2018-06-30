package main

import (
	"bufio"
	"fmt"
)

type Writer int

func (*Writer) Write(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	w := new(Writer)
	bw1 := bufio.NewWriterSize(w, 3)
	bw2 := bufio.NewWriterSize(bw1, 1)
	bw3 := bufio.NewWriterSize(bw2, 4)

	fmt.Printf("%#v\n", bw1)
	fmt.Printf("%#v\n", bw2)
	fmt.Printf("%#v\n", bw3)
}
