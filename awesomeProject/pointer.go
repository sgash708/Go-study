package main

import "fmt"

// func one(x int) {
//  // x = 1
// }

func one(x *int) {
	*x = 1 // 実態の書き換え
}

func aPoint() {

	var n int = 100
	fmt.Println(n)
	// 100を入れたメモリのアドレスが表示される
	fmt.Println(&n)

	// イントのポイント型
	var p *int = &n
	fmt.Println(p)  // メモリのアドレス
	fmt.Println(&p) // 変数「p」を格納しているアドレス
	fmt.Println(*p) // メモリのアドレスないに入っているバリューを取得

	// one(n)
	// fmt.Println(n) // 100で変化がない

	/*
		「&変数名」で送る⇨アドレス
		「(変数名 *int)」で受ける⇨実態のバリューを受けている
	*/
	one(&n)
	fmt.Println(n)
}
