package main

import "fmt"

func main() {
	// 整列された整数列
	a := []int{1, 3, 5, 9, 10, 19, 43, 77, 99, 103, 150, 320, 450, 455, 1300}
	fmt.Println("----------------array-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	fmt.Println("----------------exists-----------------")
	key := 455
	fmt.Printf("val:%s, address:%s\n", a[*binarySearch(a, key)], &a[*binarySearch(a, key)])

	fmt.Println("----------------nil-----------------")
	key = 4
	fmt.Printf("%v\n", binarySearch(a, key))
}

func binarySearch(a []int, key int) *int {
	left := 0
	right := len(a)
	for left < right {
		mid := (left + right) / 2
		if a[mid] == key {
			return &mid
		} else if a[mid] < key {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return nil
}
