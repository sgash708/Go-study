package main

import "fmt"

func aInt() {
	var (
		// golangの書き方では、「=」などは揃える
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i
	)
	fmt.Println(u8, i8, f32, c64)
	// 型確認
	// 「"%v"」でvalueを指定し、「"%T"」と合わせて出力できる
	// 「type=」「value=」は文字列として認識させられる
	fmt.Printf("type=%T value=%v\n", u8, u8)

	// 計算
	// x := 0
	// fmt.Println(x)
	// x++
	// fmt.Println(x)
	// x--
	// fmt.Println(x)

	// 二進数シフト
	// fmt.Println(1 << 0) // 0001 0001
	// fmt.Println(1 << 1) // 0001 0010
	// fmt.Println(1 << 2) // 0001 0100
	// fmt.Println(1 << 3) // 0001 1000

	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)
}