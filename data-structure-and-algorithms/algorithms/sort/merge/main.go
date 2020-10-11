package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{7, 19, 9, 12, 4, 5, 6, 14}
	fmt.Println(a)

	mergeSort(a, 0, len(a))
	fmt.Println(a)
}

func mergeSort(a []int, left, right int) {
	fmt.Printf("input left: %d, right: %d\n", left, right)

	if left+1 < right {
		mid := (left + right) / 2
		fmt.Printf("mergeSort left: %d, mid: %d\n", left, mid)
		mergeSort(a, left, mid)
		fmt.Printf("mergeSort mid: %d, right: %d\n", mid, right)
		mergeSort(a, mid, right)
		fmt.Printf("merge left: %d, mid: %d, right: %d\n", left, mid, right)
		merge(a, left, mid, right)
	}
	fmt.Println("-----end-----")
}

func merge(a []int, left, mid, right int) {
	tmp := make([]int, right)

	// 配列のコピー
	for i := left; i <= right-1; i++ {
		tmp[i] = a[i]
	}

	// 1つ目の配列は昇順に2つ目の配列は降順になるようにしている
	reverse(tmp, mid, right)

	i, j := left, right-1
	for k := left; k <= right-1; k++ {
		if tmp[i] <= tmp[j] {
			a[k] = tmp[i]
			i += 1
		} else {
			a[k] = tmp[j]
			j -= 1
		}
	}
}

func reverse(a []int, left, right int) {
	for i := left; i <= left+(right-left)/2-1; i++ {
		j := right - (i - left) - 1
		swap(&a[i], &a[j])
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
