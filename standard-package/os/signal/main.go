package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// しぐなるのしゅるい

// SIGHUP
// SIGINT
// SIGTERM
// SIGQUIT
// SIGILL
// SGTRAP
// SIGAMRT
// SIGSTKFLT
// SIGMET
// SIGSYS

func main() {
	defer fmt.Println("done")

	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	}

	// 受信するチャンネルの用意
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, trapSignals...)

	ctx, cancel := context.WithCancel(context.Background())

	go doMain(ctx)

	sig := <-sigCh
	fmt.Println("Got signal", sig)

	cancel()

	time.Sleep(3 * time.Second) // doMainのスレッドが終了したログが吐き出されるのを確認するために待つ
}

func doMain(ctx context.Context) {
	defer fmt.Println("done infinite loop")
	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done")
			return
		default:
			time.Sleep(1 * time.Second)
			fmt.Println("sleep")
		}
	}
}
