package main

import "fmt"

func aFunction() {
	// リターン①
	// added := add(10, 20)
	// fmt.Println(added)

	// リターン②
	a1, a2 := add(10, 20)
	fmt.Println(a1, a2)

	// 掛け算
	a3 := cal(10, 20)
	fmt.Println(a3)

	// 内部関数1
	f := func(x int) {
		fmt.Println("Inner func", x)
	}
	// 実行
	f(100)

	// 内部関数2
	// 関数を変数に入れ直さずにそのまま実行する方法
	// 並列化を行う
	func(x int) {
		fmt.Println("Inner func", x)
	}(1)
}

/*
	リターン①
	返り値があるときは、「()」の隣に「型」を記述する
	「func add(x, y int) int」として「xとy」を同じ型として認識させてあげることもできる
*/
// func add(x int, y int) int {
// 	return x + y
// }

/*
	リターン②
	返り値を複数個かける
*/
func add(x int, y int) (int, int) {
	return x + y, x - y
}

/*
	返り値の設定として、「result int」を作成している
	→新たに、変数名resultを作成することはできない

	「:=」で「result」を宣言するなら「(int)」にしなくてはいけない

	[メリット]
	・返り値の変数名と型名を指定しているので、「return」のみで「result」が返される
	・一行で何を返すか、わかる
*/
func cal(price, item int) (result int) {
	result = price * item
	return result
}
