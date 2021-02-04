package main

import "fmt"

// 「const」は型を定義しない
// overflowについては出力時に超えていなければ問題ない
const big = 9223372036854775807 + 1

/*
Pi 		 is 円周率
Username is ユーザ名
Password is パスワード
*/
const (
	Pi       = 3.14
	Username = "test_user"
	Password = "test_pass"
)

func learningConst() {
	fmt.Println(Pi, Username, Password, big-1)
	// 定数の値書き換えは不可能
	// Pi = 3

	// 「big」の型を確認する
	fmt.Printf("%T\n", big-1)
}
