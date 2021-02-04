package main

import (
	"fmt"
	"strings"
)

func learningString() {
	fmt.Println("Hello World!!")
	fmt.Println("Hello" + "World!!")

	// インデックス指定し出力 → ASCIIコードで表示
	fmt.Println("Hello World!!!"[0])
	// string変換
	fmt.Println(string("Hello World!!!"[0])) // 「H」が出力される

	// 文字列の代入にインデックス指定はできない
	var str string = "Hello World!"

	// 文字列の置換方法
	// string.Replace(変数名, "文字列", "置換する文字列", 置換数)
	fmt.Println(strings.Replace(str, "H", "X", 1))
	// 元の変数は変化していないこと確認
	fmt.Println(str)
	// 含まれているか(bool)
	// strings.Contains(変数名, "検索したい文字列")
	fmt.Println(strings.Contains(str, "World"))

	// 複数行にわたる文字列(バッククォートにする)
	fmt.Println(`hoge
	hoge
						hoge
hoge`)

	// 特殊文字
	fmt.Println(`"`)
	fmt.Println("\"")

	// rune型
	// シングルクォーテーションが使える
	var hoge rune = 'g'
	fmt.Printf("%T\n", hoge)
	fmt.Println(hoge)
}
