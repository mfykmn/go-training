package main

import (
	"fmt"
)

func main() {
	a := 6
	b := 3
	c := 4

	fmt.Println("----------------before-----------------")
	fmt.Printf("a\n  val:%s, address:%s\n", a, &a)
	fmt.Printf("b\n  val:%s, address:%s\n", b, &b)
	fmt.Printf("c\n  val:%s, address:%s\n", c, &c)

	sortingBySwap(&a, &b, &c)

	fmt.Println("----------------after-----------------")
	fmt.Printf("a\n  val:%s, address:%s\n", a, &a)
	fmt.Printf("b\n  val:%s, address:%s\n", b, &b)
	fmt.Printf("c\n  val:%s, address:%s\n", c, &c)
}

func sortingBySwap(a, b, c *int) {
	if *a > *b {
		swap(a, b)
	}
	if *b > *c {
		swap(b, c)
	}
	if *a > *b {
		swap(a, b)
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
