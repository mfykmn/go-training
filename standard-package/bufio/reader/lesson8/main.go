package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("abcd")
	br := bufio.NewReader(r)
	byte, err := br.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", byte)
	fmt.Printf("buffered = %d\n", br.Buffered())
	err = br.UnreadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("buffered = %d\n", br.Buffered())
	byte, err = br.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", byte)
	fmt.Printf("buffered = %d\n", br.Buffered())
}
