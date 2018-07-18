package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand" // 疑似乱数を生成するパッケージ
	"time"
)

func main() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		fmt.Println("err")
		// crypto/rand からRead出来なかった場合の代替手段
		s = time.Now().UnixNano()
	}

	rand.Seed(s)
	n := rand.Intn(100) // 0 <= n < 100 の範囲
	fmt.Println(n)
}
