package main

import "fmt"

func main() {
	a := []int{3, 5, 9, 11}
	fmt.Print(sum(a))
}

func sum(a []int) int {
	var sum int
	for _, v := range a {
		sum += v
	}
	return sum
}
