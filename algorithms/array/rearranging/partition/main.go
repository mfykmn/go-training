package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{1, 9, 12, 5, 4, 10, 6, 8}
	fmt.Println("----------------before-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	partition(a, 0, len(a)-1)
	fmt.Println("----------------after-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

}

func partition(a []int, left, right int) int {
	p := left
	i := p - 1
	for j := p; j <= right-1; j++ {
		if a[j] < a[right] {
			i += 1
			swap(&a[i], &a[j])
		}
	}

	i += 1
	swap(&a[i], &a[right])
	return i
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
