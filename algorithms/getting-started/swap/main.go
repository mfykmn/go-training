package main

import (
	"fmt"
)

func main() {
	a := 1
	b := 2

	fmt.Println("----------------before-----------------")
	fmt.Printf("a\n  val:%s, address:%s\n", a, &a)
	fmt.Printf("b\n  val:%s, address:%s\n", b, &b)

	swap(&a, &b)

	fmt.Println("----------------after-----------------")
	fmt.Printf("a\n  val:%s, address:%s\n", a, &a)
	fmt.Printf("b\n  val:%s, address:%s\n", b, &b)
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
