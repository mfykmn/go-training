package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.2
	for i := 0; i < 10; i++ {
		z = z - (math.Pow(z, 2)-x)/2*z
	}
	return z
}

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		fmt.Println(math.Sqrt(float64(i)))
		fmt.Println(Sqrt(float64(i)))
	}
}
