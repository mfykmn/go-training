package main

import "fmt"

// ラムダ式を実装している場合は以下のような形式でStrategyパターンを表すことができる
func main() {
	s := Strategy{2}
	fmt.Println(s.execute(func(in int) int {
		return in * 2
	}))

	fmt.Println(s.execute(func(in int) int {
		return in / 2
	}))
}


type Strategy struct {
	num int
}

func (s *Strategy) execute(f func (int) int) int {
	return f(s.num)
}


