package main

import "fmt"

type FanoutTask struct {
	taskFunc func()
}

func FanoutWorker(in chan FanoutTask) {
	for {
		task, ok := <-in
		if !ok {
			return
		}

		// taskを処理
		task.taskFunc()
	}

}

func FanoutDispatcher(out chan FanoutTask) {
	defer close(out)
	// データベースなどからタスクを取得していく
	for {
		task, err := FanoutGetNextTask()
		if err != nil {
			return
		}

		//channelに書き込む
		out <- task
	}
}

var count int

func FanoutGetNextTask() (FanoutTask, error) {
	count++
	if count > 20 {
		return FanoutTask{}, fmt.Errorf("無限ループになってしまうので20以上で停止")
	}
	return FanoutTask{taskFunc: func() { fmt.Println("task実行", count) }}, nil
}

func ExampleFanout() {
	ch := make(chan FanoutTask)

	for i := 0; i < 3; i++ {
		go FanoutWorker(ch)
	} // ワーカが３つ起動している

	FanoutDispatcher(ch) // ディスパッチャーが1つ起動している
}

func main() {
	ExampleFanout()
}
