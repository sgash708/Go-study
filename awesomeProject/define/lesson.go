package main

import (
	"fmt"
	"time"
)

/*
	コンストラクタ
	基本文法はセミコロンなし
	複数行：「\/**\/」
	単数行：「//」
*/

func main() {
	// fmtと記載すると勝手にインポートされる
	fmt.Println("Hello world!!", "TEST", "TEST 2", time.Now())

	// userはインポートされない
	// 理由は、「os/user」と階層が違うから
	// ドット繋ぎにせず、スラッシュで繋ぐ
	// fmt.Println(user.Current())

	// 変数について
	// hoge()
	// hugu()

	// aInt()
	// aString()
	// aBool()
	// acast()
	// aArr()
	// aSlice()
	// aSliceMakeCap()
	// aMap()
	// aByte()
	// aFunction()
	// aClo()
	// aValFunc()
	// aExer()
}
