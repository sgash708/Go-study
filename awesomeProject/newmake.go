package main

import "fmt"

func aNewMake() {
	// メモリの領域は確保している
	// アドレスは返ってくる
	var p *int = new(int)
	fmt.Println(p)  //
	fmt.Println(&p) // アドレス
	fmt.Println(*p) // 実数
	*p++
	fmt.Println(*p) // インクリメントした値

	var p2 *int
	fmt.Println(p2)
	// nilは、実数の定義をしていないのでインクリメントはできない！
	// panicを起こす
	// *p2++
	// fmt.Println(*p2)

	/*
		makeについて
		・pointerを返すか返さないかで違いが出てくる
	*/
	s := make([]int, 0)
	fmt.Printf("%T %v\n", s, s)

	smap := make(map[string]int)
	fmt.Printf("%T %v\n", smap, smap)

	ch := make(chan int)
	fmt.Printf("%T %v\n", ch, ch)

	var ps *int = new(int)
	fmt.Printf("%T %v\n", ps, ps)

	var st = new(struct{})
	fmt.Printf("%T %v\n", st, st)

}
