package main

import (
	"context"
	"fmt"
	"time"
)

func WorkWithContext(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("処理終了")
			return
		default:
			// Doneが帰ってこなければキャンセルされていない
			// ので本来の処理を行う
			fmt.Println("処理中")
		}
	}
}

func ExampleContext() {
	// キャンセル処理可能なContextを作成
	ctx, cancel :=
		context.WithCancel(context.Background())

	// この関数が終わったら必ずキャンセルされる
	defer cancel()

	// 処理を行うgoroutine
	go WorkWithContext(ctx)

	time.Sleep(3 * time.Second)

}

func main() {
	ExampleContext()
}
