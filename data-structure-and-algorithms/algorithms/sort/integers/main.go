package main

import "fmt"

func main() {
	// ランダムな列
	a := []int{3, 2, 4, 3, 4, 1}
	fmt.Println(a)
	b := countingSort(a, 4)
	fmt.Println(b)
}

func countingSort(a []int, max int) []int {
	c := make([]int, max+1, max+1)
	b := make([]int, len(a), len(a))

	for i := 0; i < len(a); i++ {
		c[a[i]]++
	}
	fmt.Println(c)

	// 一次元累積和を出す
	for i := 1; i <= max; i++ {
		c[i] = c[i] + c[i-1]
	}
	fmt.Println(c)

	for i := len(a) - 1; i >= 0; i-- {
		c[a[i]]--
		b[c[a[i]]] = a[i]
	}

	return b
}
