package main

import "fmt"

type Queue struct {
	q          []int
	head, tail int
	size       int
}

func NewQueue(size int) *Queue {
	return &Queue{
		q:    make([]int, size, size),
		head: 0,
		tail: 0,
		size: size,
	}
}

func (q *Queue) enqueue(x int) error {
	if q.tail+1 == q.head { //TODO これ本当に必要？
		return fmt.Errorf("head exceeds tail")
	}
	if q.tail > q.size-1 {
		if q.empty() {
			q.head = 0
			q.tail = 0
		} else {
			return fmt.Errorf("size over")
		}
	}

	q.q[q.tail] = x
	q.tail++
	return nil
}

func (q *Queue) dequeue() (res int, err error) {
	if q.empty() {
		return 0, fmt.Errorf("empty")
	}
	res = q.q[q.head]
	q.head++
	return
}

// emptyはheadとtailが等しいときにtrueを返す
func (q *Queue) empty() bool {
	return q.head == q.tail
}

func main() {
	queue := NewQueue(3) // 3つの変数を持ち続けられる
	fmt.Println(queue.enqueue(4))
	fmt.Println(queue.enqueue(8))
	fmt.Println(queue.enqueue(1))
	fmt.Println(queue.dequeue())  // 4
	fmt.Println(queue.enqueue(7)) // size over
	fmt.Println(queue.dequeue())  // 8
	fmt.Println(queue.dequeue())  // 1
	fmt.Println(queue.dequeue())  // empty
	fmt.Println(queue.enqueue(3))
	fmt.Println(queue.enqueue(9))
	fmt.Println(queue.dequeue()) // 3
	fmt.Println(queue.dequeue()) // 9
}
