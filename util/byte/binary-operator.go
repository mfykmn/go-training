package main

import "fmt"

func main() {
	var b byte = 0xAA // 10101010
	var m byte = 0x3C // 00111100

	fmt.Printf("b=%08b\n", b)
	fmt.Printf("m=%08b\n", m)
	fmt.Println("----------------------")

	fmt.Printf("b&m=%08b\n", b&m) // ビットごとの積 (AND)
	fmt.Printf("b|m=%08b\n", b|m) // ビットごとの和 (OR)
	fmt.Printf("b^m=%08b\n", b^m) // ビットごとの排他的論理和 (XOR)
	fmt.Printf("b&^m=%08b\n", b&^m) // ビットクリアー、論理積の否定 (AND NOT)

	//算術シフト
	c, d := 15, 240 // 00001111, 11110000
	fmt.Println("----------------------")
	fmt.Printf("b=%08b\n", c)
	fmt.Printf("m=%08b\n", d)
	fmt.Println("----------------------")

	fmt.Printf("c<<4=%08b\n", c<<4) // 11110000
	fmt.Printf("d>>4=%08b\n", d>>4) // 00001111
	fmt.Printf("b<<2=%08b\n", b<<2) // 左シフト
	fmt.Printf("b>>2=%08b\n", b>>2) // 右シフト
}

