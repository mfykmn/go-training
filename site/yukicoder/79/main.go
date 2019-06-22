package main

import (
	"fmt"
	"sort"
)

type Level struct {
	level int
	vote  int
}

type Levels []Level

func (ls Levels) Len() int {
	return len(ls)
}

func (ls Levels) Less(i, j int) bool {
	if ls[i].vote < ls[j].vote {
		return true
	} else if (ls[i].vote == ls[j].vote) && (ls[i].level < ls[j].level) {
		return true
	}
	return false
}

func (ls Levels) Swap(i, j int) {
	ls[i], ls[j] = ls[j], ls[i]
}

func main() {
	// memo 値を取り出す
	var N int
	fmt.Scan(&N)

	L := make([]int, N, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&L[i])
	}

	// memo 値のcount
	c := make(map[int]int)
	for _, v := range L {
		c[v] += 1
	}

	// memo sort
	var levels Levels
	for k, v := range c {
		levels = append(levels, Level{
			level: k,
			vote:  v,
		})
	}
	sort.Sort(sort.Reverse(levels))
	fmt.Println(levels[0].level)
}
