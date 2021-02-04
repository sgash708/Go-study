package main

import "fmt"

func learningBool() {
	t, f := true, false
	//
	fmt.Printf("%T %v %t\n", t, t, t)
	fmt.Printf("%T %v %t\n", f, f, f)

	// 間違いパターン
	// fmt.Printf("%T %v %t\n", t, 1, 1)
	// fmt.Printf("%T %v %t\n", f, 0, 0)

	// and評価
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	// or評価
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)

	// 否定
	fmt.Println(!true)
	fmt.Println(!false)
}
