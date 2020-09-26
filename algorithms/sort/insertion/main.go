package main

import "fmt"

func main() {
	// 整列された整数列
	a := []int{1, 3, 5, 9, 10, 19, 43, 77, 99, 103, 150, 320, 450, 455, 1300}
	a = append(a, 100)

	fmt.Println("----------------before-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	fmt.Println("----------------after-----------------")
	index := len(a) - 1
	insertionSort(a, index)
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}
}

func insertionSort(a []int, i int) {
	j := i - 1
	tmp := a[i]

	for {
		if j < 0 {
			break
		} else if !(j >= 0 && a[j] > tmp) {
			break
		} else {
			a[j+1] = a[j]
			j = j - 1
		}
	}

	a[j+1] = tmp
}
