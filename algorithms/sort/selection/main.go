package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{7, 9, 12, 4, 5}

	fmt.Println("----------------before-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	selectionSort(a)

	fmt.Println("----------------after-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}
}

func selectionSort(a []int) {
	n := len(a)
	for i := 0; i <= n-2; i++ {
		minj := min(a, i, n)
		swap(&a[i], &a[minj])
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}

func min(a []int, left, right int) int {
	mini := left
	for i := mini; i <= right-1; i++ {
		if a[i] < a[mini] {
			mini = i
		}
	}
	return mini
}
