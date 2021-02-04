package main

import "fmt"

func learningSlice() {
	n := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(n)
	fmt.Println(n[2])
	// 0-2_start, 0-4_end == [3, 4]
	fmt.Println(n[2:4])
	// 先頭から2番目までを取得する
	fmt.Println(n[:2])
	// 先頭から2番目より後ろを全て取得する
	fmt.Println(n[2:])
	// 全て取得
	fmt.Println(n[:])

	// 配列の書き換え
	n[2] = 100
	fmt.Println(n)

	// 多次元配列
	var boards = [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
		[]int{7, 8, 9},
	}
	fmt.Println(boards)
	for _, board := range boards {
		for _, tantai := range board {
			fmt.Println(tantai)
		}
	}

	n = append(n, 100, 200, 300, 400)
	fmt.Println(n)
}
