package main

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "ok"
	} else {
		return "ng"
	}
}

func aIf() {
	num := 9
	if num%2 == 0 {
		fmt.Println("by 2")
	} else if num%3 == 0 {
		fmt.Println("by 3")
	} else {
		fmt.Println("else")
	}

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("and演算子")
	}

	if x == 10 || y == 10 {
		fmt.Println("or演算子")
	}
	result := by2(x)
	if result == "ok" {
		fmt.Println("great")
	} else {
		fmt.Println("no good")
	}

	// 上記の処理を一行で書く
	// 「;」で処理をつなげる
	// if文の中だけで「result2」が使える
	if result2 := by2(x); result2 == "ok" {
		fmt.Println("great 2")
	}

}
