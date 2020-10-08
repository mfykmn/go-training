package main

import "fmt"

func main() {
	// 最大ヒープとなっている整数の列
	a := []int{20, 9, 11, 7, 3, 6, 9, 3, 2}
	fmt.Println(a)
	pq := NewPriorityQueue(a)
	fmt.Println("---extract---")
	fmt.Println(pq.extract())
	fmt.Println(pq.a)
	fmt.Println("---extract---")
	fmt.Println(pq.extract())
	fmt.Println(pq.a)
	fmt.Println("---insert 5---")
	pq.insert(5)
	fmt.Println(pq.a)
	fmt.Println("---top---")
	fmt.Println(pq.top())
}

func NewPriorityQueue(a []int) *PriorityQueue {
	return &PriorityQueue{
		a:        a,
		heapSize: len(a),
	}
}

type PriorityQueue struct {
	a        []int // キューの要素を保持する配列
	heapSize int   // 実際にデータを保持しているヒープサイズ
}

func (pq *PriorityQueue) insert(x int) {
	pq.heapSize++
	pq.a = append(pq.a, x)
	upHeap(pq.a, pq.heapSize-1)
}

func (pq *PriorityQueue) extract() int {
	val := pq.a[0]
	pq.a[0] = pq.a[pq.heapSize-1]
	pq.heapSize--
	downHeap(pq.a, 0)
	pq.a = pq.a[:pq.heapSize]
	return val
}

func (pq *PriorityQueue) top() int { return pq.a[0] }

func upHeap(a []int, i int) {
	for {
		if i <= 0 { // 根に到達したら終了
			break
		} else if a[i] <= a[parent(i)] { // ヒープ条件を満たしたら終了
			break
		} else { // 根に向かって移動
			swap(&a[i], &a[parent(i)])
			i = parent(i)
		}
	}
}

func downHeap(a []int, i int) {
	n := len(a) - 1
	l := left(i)
	r := right(i)
	largest := i

	if l < n && a[l] > a[i] { // 親(自分)、左の子、右の子の仲で最大のノードを見つける
		largest = l
	} else {
		largest = i
	}

	if r < n && a[r] > a[largest] {
		largest = r
	}

	if largest != i { // どちらかの子が最大の場合
		swap(&a[i], &a[largest])
		downHeap(a, largest) // 再帰によってダウンヒープを繰り返す
	}
}

func swap(a, b *int) {
	var t int
	t = *a
	*a = *b
	*b = t
}

func parent(index int) int { return (index - 1) / 2 }

func left(index int) int { return index*2 + 1 }

func right(index int) int { return index*2 + 2 }
