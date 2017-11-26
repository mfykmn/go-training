package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	arr := make([][]uint8, dy)
	for i, _ := range arr {
		arr[i] = make([]uint8, dx)
	}
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			arr[y][x] = uint8(x ^ y)
		}
	}
	return arr
}

func main() {
	pic.Show(Pic)
}
