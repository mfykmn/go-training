package main

import (
	"math/rand" // 疑似乱数を生成するパッケージ
	"fmt"
)

func main() {
	rand.Seed(42)
	n := rand.Intn(100) // 0 <= n < 100 の範囲
	fmt.Println(n) // 5 Seedが同じ限り毎回同じ結果
}
