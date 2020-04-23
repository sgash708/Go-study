package main

import (
	"fmt"
	"os"
)

func foochan() {
	defer fmt.Println("world foo") // print 2nd
	fmt.Println("Hello!! foo")     // print 1st
}
func aDefer() {
	foochan()

	// 遅延実行
	// 現在の関数の処理が終了したら一番最後に実行される処理
	defer fmt.Println("world") // print 4th
	// ↓が優先される
	fmt.Println("Hello!!") // print 3rd

	// deferは入れた順の「逆の順番」で実行される
	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")

	// 実際の使い方（本当はエラーハンドリングを行う）
	// ①ファイルを開く
	file, _ := os.Open("./lesson.go")
	// ⑤最後にファイルを閉じる
	defer file.Close()
	// ②バイト配列を作成
	data := make([]byte, 100)
	// ③ファイルを読み込む
	file.Read(data)
	// ④読み込んだバイト配列を「stringにキャスト」
	fmt.Println(string(data))
}
