package main

import "fmt"

type Stack struct {
	s   []int
	top int
}

func NewStack(size int) Stack {
	return Stack{
		s:   make([]int, size, size),
		top: -1,
	}
}

func (s *Stack) push(x int) error {
	if s.top+1 >= len(s.s) {
		return fmt.Errorf("size over")
	}
	s.top += 1
	s.s[s.top] = x
	return nil
}

func (s *Stack) pop() (res int, err error) {
	if s.top < 0 {
		return 0, fmt.Errorf("empty")
	}
	res = s.s[s.top]
	s.top -= 1
	return
}

func (s *Stack) peak() int {
	return s.s[s.top]
}

func (s *Stack) empty() int {
	s.top = -1
	return s.top
}

func (s *Stack) size() int {
	return s.top + 1
}

func main() {
	st := NewStack(2) // 配列の大きさを2にしておく
	fmt.Println(st.push(8))
	fmt.Println(st.push(6))
	fmt.Println(st.push(7)) // size over
	fmt.Println(st.pop())
	fmt.Println(st.push(5))
	fmt.Println(st.pop())
	fmt.Println(st.pop())
	fmt.Println(st.push(1))
	fmt.Println(st.peak())
	fmt.Println(st.empty())
	fmt.Println(st.pop()) // empty
}
