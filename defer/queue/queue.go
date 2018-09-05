package queue

import "sync"



type queue struct {

}
func New() *queue {
	return &queue{}
}

func (j *queue) Dequeue() *Job {
	return &Job{}
}

type Job struct {
	mu sync.Mutex // 排他制御
	released bool // 処理済みかどうか
	success bool // 正常終了かどうか
}

func (j *Job) Release() {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.released { // すでにリリースされている
		return
	}
	if j.success {
		// 正常終了。分散キューから削除する
	} else {
		// 異常終了。分散キューにもう一度載せる
	}
}

func (j *Job) Done() {
	j.mu.Lock()
	defer j.mu.Unlock()
	if j.released {
		return
	}
	j.success = true
}