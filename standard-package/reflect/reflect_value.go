package main

import (
	"reflect"
	"fmt"
)

func main() {
	rv1 := reflect.ValueOf(1)
	//rv2 := reflect.ValueOf("Hello World")
	//rv3 := reflect.ValueOf([]byte{0xa,0xd})
	//rv4 := reflect.ValueOf(make(chan struct{}))

	fmt.Println(rv1.Int())
	//fmt.Println(rv2.Int()) #panic
	//fmt.Println(rv3.Int()) #panic
	//fmt.Println(rv4.Int()) #panic
}
