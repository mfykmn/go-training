package main

import (
	"bufio"
	"fmt"
)

type Writer1 int

func (*Writer1) Write(p []byte) (n int, err error) {
	fmt.Printf("writer#1: %q\n", p)
	return len(p), nil
}

type Writer2 int

func (*Writer2) Write(p []byte) (n int, err error) {
	fmt.Printf("writer#2: %q\n", p)
	return len(p), nil
}

// このプログラムには1つのバグがあります。Resetを呼び出す前に、Flushでバッファをフラッシュする必要があります
// 動き的には構造体の中身を更新してくれるだけ。効果としては同じメモリ空間を使うので、作り直しの場合はGCが走るけどResetは走らないっていう利点がある
func main() {
	w1 := new(Writer1)
	bw := bufio.NewWriterSize(w1, 2)
	bw.Write([]byte("ab"))
	bw.Write([]byte("cd"))
	w2 := new(Writer2)
	bw.Reset(w2)
	fmt.Printf("%#v", bw)
	bw.Flush()
	bw.Write([]byte("ef"))
	bw.Flush()
}
