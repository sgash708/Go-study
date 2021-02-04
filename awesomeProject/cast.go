package main

import (
	"fmt"
	"strconv"
)

func learningCast() {
	// float型変換
	var x int = 1
	xx := float64(x)
	fmt.Printf("%T %v %f \n", xx, xx, xx)

	// int変換
	var y float64 = 1.2
	yy := int(y)
	fmt.Printf("%T %v %d \n", yy, yy, yy)

	// string変換
	var s string = "14"
	// string型をint型にキャストできない
	// strconvパッケージを使い「Atoi」を取得
	// Atoiはstringとerrorを返す
	// そのため「i」「_」を使用しているが、変数として出力しなくても問題はない
	i, _ := strconv.Atoi(s)
	fmt.Printf("%T %v\n", i, i)

	// byte型であれば、string型にキャストできる
	h := "Hello World!!"
	fmt.Println(string(h[0]))
}
