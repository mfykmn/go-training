package main

import "fmt"

func main() {
	fmt.Print(gcd(8, 12))
}

func gcd(a, b int) int {
	for 0 < b {
		r := a % b
		a = b
		b = r
	}
	return a
}
