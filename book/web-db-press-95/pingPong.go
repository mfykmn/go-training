package main

import (
	"fmt"
	"sync"
)

type PingPongPayload struct {
	Counter int
}

func ExamplePingPong() {
	var p PingPongPayload
	chA := make(chan *PingPongPayload)
	chB := make(chan *PingPongPayload)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			p, ok := <-chA // chAから読み込み
			if !ok {
				break
			}
			fmt.Printf("chA: p.Counter = %d\n", p.Counter)
			p.Counter++ // データの内容を書き換え
			if p.Counter > 6 {
				break
			}

			chB <- p // chBへ書き込み
		}

		close(chB)
	}()

	go func() {
		defer wg.Done()
		for {
			p, ok := <-chB // chBから読み込み
			if !ok {
				break
			}
			fmt.Printf("chB: p.Counter = %d\n", p.Counter)
			p.Counter++ // データの置き換え
			if p.Counter > 6 {
				break
			}

			chA <- p // chAへ書き込み
		}
		close(chA)
	}()

	chA <- &p // chAへの最初の書き込み
	wg.Wait() // 終了を待つ
}

func main() {
	ExamplePingPong()
}
