package main

import "fmt"

func main() {
	// ヒープ条件を満たす整数の列
	a := []int{20, 18, 11, 5, 7, 6, 9, 3, 2, 4, 2, 5}
	fmt.Println(a)
	decrease(a, 0, 1)
	fmt.Println(a)
	downHeap(a, 0)
	fmt.Println(a)

}

func decrease(a []int, i, value int) {
	a[i] = value
}

func downHeap(a []int, i int) {
	n := len(a) - 1
	l := left(i)
	r := right(i)
	largest := i

	if l < n && a[l] > a[i] { // 親(自分)、左の子、右の子の仲で最大のノードを見つける
		largest = l
	} else {
		largest = i
	}

	if r < n && a[r] > a[largest] {
		largest = r
	}

	if largest != i { // どちらかの子が最大の場合
		swap(&a[i], &a[largest])
		downHeap(a, largest) // 再帰によってダウンヒープを繰り返す
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}

func left(n int) int {
	return n*2 + 1
}

func right(n int) int {
	return n*2 + 2
}
