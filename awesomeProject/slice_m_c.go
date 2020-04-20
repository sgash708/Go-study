package main

import "fmt"

func aSliceMakeCap() {
	// sliceの数の初期値は「3」
	// sliceの数の最大値は「5」
	n := make([]int, 3, 5)
	fmt.Printf("len=%d cap=%d val=%v\n", len(n), cap(n), n)

	// 追加
	n = append(n, 0, 0)
	fmt.Printf("len=%d cap=%d val=%v\n", len(n), cap(n), n)

	// もし最大値を超えていても「増やすこと」が可能
	n = append(n, 1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("len=%d cap=%d val=%v\n", len(n), cap(n), n)

	// makeの引数を一つにした場合
	// 「長さ」と「最大数」は同一になる
	a := make([]int, 3)
	fmt.Printf("len=%d cap=%d val=%v\n", len(a), cap(a), a)

	b := make([]int, 0)
	// 宣言だけして中身を作っていない
	var c []int
	fmt.Printf("len=%d cap=%d val=%v\n", len(b), cap(b), b)
	fmt.Printf("len=%d cap=%d val=%v\n", len(c), cap(c), c)

	// 実行結果が異なる
	// 「追記」
	c = make([]int, 5) // [0 0 0 0 0 1 2 3 4]
	// 「上書き」
	// c = make([]int, 0, 5) // [0 1 2 3 4]
	for i := 0; i < 5; i++ {
		c = append(c, i)
		fmt.Println(c)
	}
	fmt.Println(c)
}
