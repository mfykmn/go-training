package main

import "fmt"

func main() {
	// ランダムな整数列
	a := []int{7, 9, 12, 4, 5}

	fmt.Println("----------------before-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}

	insertionSort(a)

	fmt.Println("----------------after-----------------")
	for i, v := range a {
		fmt.Printf("val:%s, address:%s\n", v, &a[i])
	}
}

func insertionSort(a []int) {
	n := len(a)
	for i := 1; i <= n-1; i++ {
		insertion(a, i)
	}
}

func insertion(a []int, i int) {
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
