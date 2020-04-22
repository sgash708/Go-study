package main

import "fmt"

/*
	変数xをメインで作成すると、他の関数で書き換える可能性がある
	→そのため他の関数として切り出す

	・返り値に、「func() int型」を返す
	→そのためfunc()をリターンする(line.15)
*/
func incrementGenerator() func() int {

	x := 0
	return func() int {
		x++
		return x
	}
}

/*
	使い方）
	①circleAreaの引数に変数piを設定
	→変数に入れ替え
	②入れ替えた変数を使い、無名関数の引数にradiusを入れる
	③返り値をメイン関数内で出力

	解説）
	円周率の定義が変わってしまっても対応できる

*/
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		return pi * radius * radius
	}
}

func aClo() {
	// 変数にファンクションを入れている場合
	// 使い方は「変数名()」で起動する
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	// checkValue()

	c1 := circleArea(3.14)
	fmt.Println(c1(10))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}

// func checkValue() {
// 	// 変数名xは呼び出すことができない
// 	// [理由]incrementGenerator()のみで変数xを使えるようにしているため
// 	// →未定義の状態となる
// 	fmt.Println(x)
// }
