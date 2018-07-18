package main

import (
	"math/rand" // 疑似乱数を生成するパッケージ
	"fmt"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(100) // 0 <= n < 100 の範囲
	fmt.Println(n)
}
