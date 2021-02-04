package main

import (
	"fmt"
	"os/user"
	"time"
)

/*
	コンストラクタ
	基本文法はセミコロンなし
	複数行：「\/**\/」
	単数行：「//」
	func init() {
		fmt.Println("init!!")
	}
*/

func main() {
	// fmtと記載すると勝手にインポートされる
	fmt.Println("Hello world!!", time.Now().Format("2006-01-02"))

	// userはインポートされない
	// 理由は、「os/user」と階層が違うから
	// ドット繋ぎにせず、スラッシュで繋ぐ
	fmt.Println(user.Current())

	// ★セクション３
	learningVal()
	learningConst()
	learningNum()
	learningString()
	learningBool()
	learningCast()
	learningArr()
	learningSlice()
	// aSliceMakeCap()
	// aMap()
	// aByte()
	// aFunction()
	// aClo()
	// aValFunc()
	// aExer()

	// ★セクション４
	// aIf()
	// aFor()
	// aRange()
	// aSw()
	// aDefer()
	// aLog()
	// aTry()
	// aPanicRecover()
	// aExer2()

	// セクション５
	// aPoint()
	// aNewMake()
	// aStruct()
	// aExer3()

	// セクション６

}
