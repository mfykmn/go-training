package main

import "fmt"

func main() {
	n := 10

	rects := []Rect{
		{x1: 1, y1: 1, x2: 5, y2: 5},
		{x1: 4, y1: 6, x2: 9, y2: 8},
	}

	res := TwoDimensionalAccumulation(n, rects)
	fmt.Print(res)
}

type Rect struct {
	x1, y1, x2, y2 int
}

func TwoDimensionalAccumulation(n int, rects []Rect) [][]int {
	var ac [][]int
	for _, v := range rects {
		x1 := v.x1
		y1 := v.y1
		x2 := v.x2
		y2 := v.y2

		ac[x1][y1]++
		ac[x2][y2]++
		ac[x1][y2]--
		ac[x2][y1]--
	}

	// 水平方向の累積和
	for x := 1; x <= n-1; x++ {
		for y := 1; y <= n-1; y++ {
			ac[x][y] += ac[x-1][y]
		}
	}
	// 垂直方向の累積和
	for y := 1; y <= n-1; y++ {
		for x := 1; x <= n-1; x++ {
			ac[x][y] += ac[x][y-1]
		}
	}

	return ac
}
