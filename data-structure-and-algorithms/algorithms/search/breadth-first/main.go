package main

import "fmt"

// TODO: 結果がうまいことでていない
func main() {
	var g Undirected
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(1, 5)
	g.AddEdge(1, 2)
	g.AddEdge(2, 5)
	g.AddEdge(2, 6)
	g.AddEdge(2, 4)
	g.AddEdge(4, 6)

	if err := breadthFirstSearch(&g, 0, 7); err != nil {
		fmt.Println(err)
	}
}

func breadthFirstSearch(graph *Undirected, start, size int) error {
	queue := NewQueue(size)
	color := make([]string, queue.size, queue.size)

	for i := 0; i <= queue.size-1; i++ {
		color[i] = "White"
	}

	fmt.Printf("init\n%#v\n", color)

	color[start] = "Gray"
	queue.enqueue(start)

	for {
		if queue.empty() {
			return nil
		}
		u, err := queue.dequeue()
		if err != nil {
			return err
		}
		color[u] = "Black"
		for _, v := range graph.AdjacencyList[u] {
			if color[v] == "White" {
				color[v] = "Gray"
				queue.dequeue()
			}
			fmt.Printf("%#v\n", color)
		}
		fmt.Printf("end\n%#v\n", color)
	}
}

type AdjacencyList [][]int

type Undirected struct {
	AdjacencyList
}

func (p *Undirected) AddEdge(n1, n2 int) {
	// Similar code in LabeledAdjacencyList.AddEdge.

	// determine max of the two end points
	max := n1
	if n2 > max {
		max = n2
	}
	// expand graph if needed, to include both
	g := p.AdjacencyList
	if int(max) >= len(g) {
		p.AdjacencyList = make(AdjacencyList, max+1)
		copy(p.AdjacencyList, g)
		g = p.AdjacencyList
	}
	// create one half-arc,
	g[n1] = append(g[n1], n2)
	// and except for loops, create the reciprocal
	if n1 != n2 {
		g[n2] = append(g[n2], n1)
	}
}

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
