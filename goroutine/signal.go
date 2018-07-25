package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	// サイズが1より大きいチャネル作成
	signals := make(chan os.Signal, 1)
	// SIGINT (Ctrl+C)を受け取る
	signal.Notify(signals, syscall.SIGINT)

	// シグナルがくるまで待つ
	fmt.Println("waiting SIGINT (CTRL+C)")
	<-signals
	fmt.Println("SIGINT arrived")
}
