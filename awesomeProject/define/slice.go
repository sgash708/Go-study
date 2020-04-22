package main

import "fmt"

func aSlice() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	// nという配列のインデックス2を使用する
	fmt.Println(n[2])

	// n配列のインデックス2から、先頭から４番目までを取得する
	fmt.Println(n[0:4])
	// 何も書かないと先頭から2番目までを取得する
	fmt.Println(n[:2])
	// 先頭から2番目より後ろを全て取得する
	fmt.Println(n[2:])

	// 配列の書き換え
	n[2] = 100
	fmt.Println(n)

	// 多次元配列
	var board = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(board)

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)
}
