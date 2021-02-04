package main

import "fmt"

// 関数外でも使用可能
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	xi = 2
	// float64 → float32へ型を変えたい場合
	var xf32 float32 = 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)

}

func learningVal() {
	// 初期値を入れないとdefaultが与えられる
	fmt.Println(i, f64, s, t, f)
	// Printf練習
	foo()
}
