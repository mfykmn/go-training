package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s1 := strings.NewReader("abcd")
	r := bufio.NewReader(s1)
	b := make([]byte, 3)
	_, err := r.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
	s2 := strings.NewReader("efgh")
	r.Reset(s2)
	_, err = r.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%q\n", b)
}
