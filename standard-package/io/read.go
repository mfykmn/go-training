package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("--ReadAll--")
	reader := strings.NewReader("Hello, world!\n")
	buffer, _ := ioutil.ReadAll(reader)
	fmt.Println(string(buffer))

	fmt.Println("--ReadFull--")
	reader1 := strings.NewReader("Hello, world!\n")
	buffer1 := make([]byte, 4)
	size, _ := io.ReadFull(reader1, buffer1)
	fmt.Println(string(buffer1))
	fmt.Println(size)

	size1, _ := io.ReadFull(reader1, buffer1)
	fmt.Println(string(buffer1))
	fmt.Println(size1)

}
