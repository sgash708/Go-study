package main

import "fmt"

// 関数外でも使用可能
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "hoge"
	t, f bool    = true, false
)

func foo() {

	// short関数は関数外では使用不可能
	ai := 1
	// 変数の中身書き換え
	ai = 2

	// float64→float32へ型を変えたい場合
	// af64 := 1.2
	var af64 float32 = 1.2

	as := "test"
	at, af := true, false
	fmt.Println(ai, af64, as, at, af)
	// 変数の方が何なのか調べてくれる
	fmt.Printf("%T\n", af64)
	fmt.Printf("%T\n", ai)
}

func hoge() {

	// 初期値を入れないとdefaultが与えられる
	// var (
	// 	i    int
	// 	f64  float64
	// 	s    string
	// 	t, f bool
	// )
	fmt.Println(i, f64, s, t, f)
	foo()
}
