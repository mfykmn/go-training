package main

import (
	"reflect"
	"fmt"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := &Point{X: 10, Y: 15}
	rv := reflect.ValueOf(p).Elem()

	fmt.Printf("rv.Type = %v\n", rv.Type())
	fmt.Printf("rv.Kind = %v\n", rv.Kind())
	fmt.Printf("rv.Interface = %v\n", rv.Interface())

	xv := rv.Field(0)
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after setInt) = %d\n", xv.Int())
}
