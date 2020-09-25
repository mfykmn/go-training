package main

import "fmt"

func main() {
	a := []int{5, 3, 2, 2, 9}
	fmt.Print(minv(a))
}

func minv(a []int) (minv int) {
	minv = a[0]
	for _, v := range a {
		if v < minv {
			minv = v
		}
	}
	return
}
