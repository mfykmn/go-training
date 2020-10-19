package main

import "fmt"

func main() {
	a := []int{1, 2, 4, 5, 8, 9, 3}
	fmt.Println(a)

	shellSort(a, len(a))
	fmt.Println(a)
}

func shellSort(a []int, n int) {
	interval := []int{5, 3, 1}
	for _, g := range interval {
		insertionSort(a, n, g)
	}
}

func insertionSort(a []int, n, g int) {
	for i := g; i <= n-1; i++ {
		insertion(a, i)
	}
}

func insertion(a []int, i int) {
	j := i - 1
	tmp := a[i]

	for {
		if j < 0 {
			break
		} else if !(j >= 0 && a[j] > tmp) {
			break
		} else {
			a[j+1] = a[j]
			j = j - 1
		}
	}

	a[j+1] = tmp
}
