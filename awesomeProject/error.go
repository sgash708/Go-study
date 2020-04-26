package main

import (
	"fmt"
	"log"
	"os"
)

func aTry() {
	file, err := os.Open("./lesson.go")
	if err != nil {
		log.Fatalln("error")
	}
	// 必ずファイルを閉じること
	defer file.Close()

	// バイト配列を作成
	data := make([]byte, 100)
	// バイト配列に変換したものを読み込ませる
	// 返り値で「読み込んだ文字数」「エラー」が戻ってくる
	// 変数「err」を二回使っているが、新しく上書きしているので問題はない
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("error Read")
	}
	// 文字型に変換して出力
	fmt.Println(count, string(data))

	// イニシャライズしてあるので、上書きしかできない
	// err = os.Chdir("test")
	// if err != nil {
	// 	log.Fatalln("error chdir")
	// }

	// 'エラーしか戻らない時'には、一行で書く場合もある
	// 記述：if 行いたい処理; err != nil {}
	if err = os.Chdir("test"); err != nil {
		log.Fatalln("error chdir")
	}
}
