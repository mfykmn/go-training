package main

import (
	"strings"
	"fmt"
)

var source2 = "123 1.234 1.0e4 test"

func main() {
	reader := strings.NewReader(source2)
	var i int
	var f, g float64
	var s string
	// Fscanの区切り文字は スペース、Fscanfの区切り文字は カンマ+スペース
	fmt.Fscan(reader, &i, &f, &g, &s)
	fmt.Printf("i=%v f=%v g=%v s=%v\n", i, f, g, s)

}
