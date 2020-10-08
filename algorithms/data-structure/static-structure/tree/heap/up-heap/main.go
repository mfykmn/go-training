package main

import "fmt"

func main() {
	// ヒープ条件を満たす整数の列
	a := []int{23, 18, 11, 7, 5, 6, 9, 3, 2, 1, 2, 3}
	fmt.Println(a)
	increase(a, 9, 25)
	fmt.Println(a)
	upHeap(a, 9)
	fmt.Println(a)

}

func increase(a []int, i, value int) {
	a[i] = value
}

func upHeap(a []int, i int) {
	for {
		if i <= 0 { // 根に到達したら終了
			break
		} else if a[i] <= a[parent(i)] { // ヒープ条件を満たしたら終了
			break
		} else { // 根に向かって移動
			swap(&a[i], &a[parent(i)])
			i = parent(i)
		}
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}

func parent(n int) int {
	return (n - 1) / 2
}
