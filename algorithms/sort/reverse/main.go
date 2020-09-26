package main

import "fmt"

func main() {
	// 整列された整数列
	a := []int{1, 3, 5, 9, 10, 19, 43, 77, 99, 103, 150, 320, 450, 455, 1300}
	fmt.Println("----------------before-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	fmt.Println("----------------after-----------------")
	reverse(a)
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}
}

func reverse(a []int) {
	n := len(a)
	for i := 0; i <= n/2-1; i++ {
		j := n - i - 1
		swap(&a[i], &a[j])
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}
