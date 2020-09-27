package main

import "fmt"

func main() {
	a := []int{5, 3, 2, 2, 9}
	fmt.Print(a[min(a)])
}

func min(a []int) (min int) {
	for i, v := range a {
		if v < a[min] {
			min = i
		}
	}
	return
}
