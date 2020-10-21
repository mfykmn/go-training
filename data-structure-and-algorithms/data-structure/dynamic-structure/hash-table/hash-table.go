package main

import "fmt"

func main() {
	hashTable := &HashTable{
		N:   5,
		Key: make([]int, 5, 5),
	}
	hashTable.insert(4)
	hashTable.insert(10)
	hashTable.insert(3)
	hashTable.insert(2)
	hashTable.insert(9)
	fmt.Println(hashTable.Key)
}

type HashTable struct {
	N   int
	Key []int
}

func (h *HashTable) h1(k int) int {
	return k % h.N
}

func (h *HashTable) h2(k int) int {
	return 1 + (k % (h.N - 1))
}

func (h *HashTable) hash(k, i int) int {
	return (h.h1(k) + i*h.h2(k)) % h.N
}

func (h *HashTable) insert(k int) int {
	i := 0 // 賞取る回数
	for {
		pos := h.hash(k, i)
		if h.Key[pos] == 0 {
			h.Key[pos] = k
			return pos // 場所を返して終了
		} else {
			// 空いていない場合は衝突回数を加算して再試行
			i++
		}
	}
}
