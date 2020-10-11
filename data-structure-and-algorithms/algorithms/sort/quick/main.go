package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{7, 19, 9, 12, 4, 5, 6, 14}
	fmt.Println(a)

	quickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func quickSort(a []int, left, right int) {
	fmt.Printf("input left: %d, right: %d\n", left, right)

	if left < right {
		fmt.Printf("partition left: %d, right: %d\n", left, right)
		q := partition(a, left, right)

		r := q - 1
		fmt.Printf("quickSort left: %d, right: %d\n", left, r)
		quickSort(a, left, r)

		l := q + 1
		fmt.Printf("quickSort left: %d, right: %d\n", l, right)
		quickSort(a, l, right)

	}
	fmt.Println("-----end-----")
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
