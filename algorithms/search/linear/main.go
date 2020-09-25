package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{3, 5, 2, 10, 1}
	fmt.Println("----------------array-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	fmt.Println("----------------exists-----------------")
	key := 5
	fmt.Printf("val:%s, address:%s\n", a[*linearSearch(&a, key)], &a[*linearSearch(&a, key)])

	fmt.Println("----------------nil-----------------")
	key = 4
	fmt.Printf("%v\n", linearSearch(&a, key))
}

// linearSearchはランダムな整数列に対する探索
func linearSearch(a *[]int, key int) *int {
	for i, v := range *a {
		if v == key {
			return &i
		}
	}
	return nil
}
