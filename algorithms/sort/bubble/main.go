package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{7, 9, 12, 4, 5}

	fmt.Println("----------------before-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	bubbleSort(a)

	fmt.Println("----------------after-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}
}

func bubbleSort(a []int) {
	n := len(a)
	for i := 0; i <= n-2; i++ {
		for j := n - 1; j >= i+1; j-- {
			if a[j-1] > a[j] {
				swap(&a[j-1], &a[j])
			}
		}
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
