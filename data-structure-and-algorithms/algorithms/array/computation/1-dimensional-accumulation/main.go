package main

import "fmt"

func main() {
	n := 10

	segments := []Segment{
		{begin: 3, end: 5},
		{begin: 2, end: 9},
		{begin: 1, end: 4},
	}

	res := OneDimensionalAccumulation(n, segments)
	fmt.Print(res)
}

type Segment struct {
	begin, end int
}

func OneDimensionalAccumulation(n int, segments []Segment) []int {
	ac := make([]int, n, n)
	for _, v := range segments {
		b := v.begin
		e := v.end
		ac[b]++
		ac[e]--
	}

	for i := 1; i <= n-1; i++ {
		ac[i] += ac[i-1]
	}
	return ac
}
