package main

import (
	"fmt"
	"strconv"
)

func acast() {
	// float型変換
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f \n", xx, xx, xx)

	// int変換
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d \n", yy, yy, yy)

	// string変換
	var sss string = "14"
	// string型をint型にキャストできない
	// ss := int(s)
	// fmt.Printf("%T %v %d", ss, ss, ss)

	// strconvパッケージを使い「Atoi」を取得
	// Atoiはstringとerrorを返す
	// そのため今回は、「i」「_」を使用している
	// 尚、「_」は変数として出力しなくても問題はない
	i, _ := strconv.Atoi(sss)
	fmt.Printf("%T %v\n", i, i)

	// byte型であれば、string型にキャストできる
	h := "Hello World!!"
	fmt.Println(string(h[0]))
}
