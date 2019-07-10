package iterator

import "fmt"

type Iterator interface {
	Iterate() (int, bool)
}


type ChanIterator struct {
	ch chan int
}

func NewIterator(data []int) *ChanIterator {
	ch := make(chan int, len(data))
	for _, v := range data {
		ch<-v
	}
	close(ch)
	return &ChanIterator{ch: ch}
}

func (iter *ChanIterator) Iterate() (int, bool) {
	v, ok := <-iter.ch
	return v, ok
}

func main() {
	iter := NewIterator([]int{1, 2, 3, 4, 5})
	for {
		v, ok := iter.Iterate()
		if !ok {
			break
		}
		fmt.Printf("%d\n", v)
	}
}

