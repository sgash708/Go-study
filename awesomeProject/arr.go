package main

import (
	"fmt"
)

func aArr() {
	var arr [2]int
	arr[0] = 100
	arr[1] = 200
	// 配列でもprintlnできる
	fmt.Println(arr)
	fmt.Println(arr[0], arr[1])

	// 初期化もできる
	// 配列は「append」できない
	var b [2]int = [2]int{100, 200}
	fmt.Println(b)

	// スライス（C#やJavaでいうリスト）も作れる
	var c []int = []int{200, 300}
	c = append(c, 400)
	fmt.Println(c)
}
