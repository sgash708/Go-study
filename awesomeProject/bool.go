package main

import "fmt"

func aBool() {
	// var t,f bool = true, false
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, t)
	fmt.Printf("%T %v %t\n", f, f, f)

	// 間違いパターン
	// fmt.Printf("%T %v %t\n", t, 1, 1)
	// fmt.Printf("%T %v %t\n", f, 0, 0)

	// &&評価
	// 冗長と出るけどテストコードなので気にしない
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)

	// ||評価
	// 冗長と出るけどテストコードなので気にしない
	fmt.Println(true || true)   // true
	fmt.Println(true || false)  // true
	fmt.Println(false || false) // false

	// 否定
	fmt.Println(!true)
	fmt.Println(!false)
}
