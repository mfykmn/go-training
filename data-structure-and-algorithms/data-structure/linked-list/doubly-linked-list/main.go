package main

import "fmt"

func main() {
	list := New()
	list.insert(1)
	list.insert(2)
	list.insert(3)
	list.insert(4)
	list.deleteKey(3)
	fmt.Printf("%#v", list.listSearch(2))
}

type Node struct {
	prev, next *Node
	key        int
}

type LinkedList struct {
	sentinel *Node
}

func New() *LinkedList {
	sentinel := &Node{}
	sentinel.prev = sentinel
	sentinel.next = sentinel
	return &LinkedList{sentinel: sentinel}
}

func (ll *LinkedList) insert(data int) {
	node := &Node{
		key:  data,
		next: ll.sentinel.next,
		prev: ll.sentinel,
	}

	ll.sentinel.next.prev = node
	ll.sentinel.next = node
}

func (ll *LinkedList) listSearch(k int) *Node {
	current := ll.sentinel.next
	for current != ll.sentinel && current.key != k {
		current = current.next
	}
	return current
}

func (ll *LinkedList) deleteNode(target *Node) {
	if target == ll.sentinel {
		return
	}

	target.prev.next = target.next
	target.next.prev = target.prev

	target.next, target.prev = nil, nil // avoid memory leaks
}

func (ll *LinkedList) deleteKey(key int) {
	ll.deleteNode(ll.listSearch(key))
}
