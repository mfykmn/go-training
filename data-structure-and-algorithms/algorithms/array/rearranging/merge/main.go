package main

import "fmt"

func main() {
	// 整列された整数列
	a := []int{1, 3, 5, 9, 10, 19, 43, 77, 99, 103, 150, 320, 450, 455, 1300}

	// 整列された整数列
	b := []int{2, 34, 100, 300}

	list := append(a, b...)
	fmt.Println("----------------before-----------------")
	for i, v := range list {
		fmt.Printf("val:%s, address:%s\n", v, &list[i])
	}
	left := 0
	mid := len(a)
	right := mid + len(b)

	merge(list, left, mid, right)
	fmt.Println("----------------after-----------------")
	for i, v := range list {
		fmt.Printf("val:%s, address:%s\n", v, &list[i])
	}
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
