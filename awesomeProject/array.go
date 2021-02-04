package main

import (
	"fmt"
)

func learningArr() {
	var arr [2]int
	arr[0] = 100
	arr[1] = 200
	fmt.Println(arr)

	// 「配列」は要素追加できない
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	/*
	 * 番外編
	 * 「スライス」は要素追加できる
	 */
	var c []int = []int{200, 300}
	c = append(c, 400)
	fmt.Println(c)
}
