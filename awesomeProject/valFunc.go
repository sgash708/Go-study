package main

import "fmt"

/*
	可変長変数について
	使い方）「...」で可変長にすることができる
*/
func aValFunc() {
	// 可変長引数なので何も渡さないということもできる
	fofo()
	fofo(10, 20)
	fofo(10, 20, 30)

	// 配列を入れたい時には「...」で引数として渡せる
	s := []int{1, 2, 3}
	fofo(s...)
}

func fofo(params ...int) {
	fmt.Println(len(params), params)

	// foreach文
	// ショートカット => 「forr」
	// 「_」はキー、「param」はバリュー
	for _, param := range params {
		fmt.Println(param)
	}
}
