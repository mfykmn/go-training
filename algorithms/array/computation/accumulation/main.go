package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{0, 19, 4, 1, 39, 5, 58, 23, 9} // 配列の１つ目は必ず0にする
	n := len(a)
	fmt.Println(accumulation(a, n, 3, 1))
}

func accumulation(a []int, n, right, left int) int {
	ac := make([]int, n, n)
	ac[0] = 0

	for i, v := range a {
		if i == 0 {
			continue
		}
		ac[i] = ac[i-1] + v
	}

	return ac[right] - ac[left-1]
}
