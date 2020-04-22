package main

import (
	"fmt"
	"strings"
)

func aString() {
	fmt.Println("Hello World!!")
	fmt.Println("string" + "add string")

	// インデックス指定し出力することもできる
	// 下記コードでは、ASCIIコードで表示
	fmt.Println("Hello World!!!"[0])

	// phpでいう「substr()」にあたる
	// string変換
	fmt.Println(string("Hello World!!!"[0])) // 「H」が出力される

	// 文字列の代入にインデックス指定はできない
	var str string = "Hello World!"
	// str[0] = "X"

	// 文字列の置換方法
	// phpでいうstr_replace()にあたる
	// string.Replace(変数名, "文字列", "置換する文字列", 置換数)
	// stringsパッケージは自動インポートされる
	fmt.Println(strings.Replace(str, "H", "X", 1))

	// 置換
	repStr := strings.Replace(str, "H", "X", 1)
	// 含まれているか(bool)
	// strings.Contains(変数名, "検索したい文字列")
	fmt.Println(strings.Contains(repStr, "World"))
	// 元の変数は変化していないこと確認
	fmt.Println(str)

	// 複数行にわたる文字列
	fmt.Println(`hoge
	hoge
						hoge
hoge`)

	// 特殊文字
	fmt.Println(`"`)

	// rune型
	// シングルクォーテーションが使える
	var hoge rune = 'g'
	fmt.Printf("%T\n", hoge)
	fmt.Println(hoge)
}
