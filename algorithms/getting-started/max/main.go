package main

import "fmt"

func main() {
	a := 9
	b := 7
	fmt.Print(max(a, b))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
