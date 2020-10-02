package main

import (
	"fmt"
	"math"
)

func main() {
	a := make([]int, 101, 101)

	sieveOfEratosthenes(a)

	fmt.Printf("-------------------\n")
	for i, v := range a {
		if i == 0 {
			continue
		}
		fmt.Printf("[%v]%v", i, v)
		if i%10 == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Printf("-------------------\n")
	for i, v := range a {
		if v == 1 {
			fmt.Printf("%v,", i)
		}
	}
}

func sieveOfEratosthenes(a []int) {
	n := len(a)

	// indexが0と１以外に1の値を入れる
	for i := 2; i <= n-1; i++ {
		a[i] = 1
	}

	// 配列長の平方根取得
	squareRoot := int(math.Sqrt(float64(n - 1)))
	for i := 2; i <= squareRoot; i++ {
		for j := 2; i*j <= n-1; j++ {
			if a[i] == 0 {
				continue
			}
			a[i*j] = 0
		}
	}
}
