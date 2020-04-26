package main

import "fmt"

func aExer2() {
	// 問題１
	// 惜しかった
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	// 自分のこたえ
	// hikaku := 1000

	// 答え：宣言だけしとく
	var min int
	for i, val := range l {
		if i == 0 {
			min = val
			continue
		}
		if min >= val {
			min = val
		}
	}
	fmt.Println(min)

	// 問題２
	map2 := map[string]int{
		"apple":  200,
		"banana": 300,
		"orange": 150,
		"grapes": 80,
		"papaya": 500,
		"kiwi":   90,
	}
	sum := 0
	for _, val := range map2 {
		sum += val
	}
	fmt.Println(sum)
}
