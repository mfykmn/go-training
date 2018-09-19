package main

import "github.com/mafuyuk/go-training/standard-package/deferard-package/defer/queue"

// learn web+db press vol.106
func main() {
	q := queue.New()
	job := q.Dequeue()
	defer job.Release()

	if err := process(job); err != nil {
		return
	}
	job.Done()
}
