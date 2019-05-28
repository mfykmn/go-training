package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	arr := map[string]int{}
	for _, v := range strings.Fields(s) {
		arr[v] += 1
	}
	return arr
}

func main() {
	wc.Test(WordCount)
}
