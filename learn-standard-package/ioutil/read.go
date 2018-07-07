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
	buffer, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buffer))

	fmt.Println("--ReadFull--")
	reader1 := strings.NewReader("Hello, world!\n")
	buffer1 := make([]byte, 4)
	size, err := io.ReadFull(reader1, buffer1)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(buffer1))
	fmt.Println(size)

}
