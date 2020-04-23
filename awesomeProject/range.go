package main

import "fmt"

func aRange() {

	// 配列の取り出し
	l := []string{"python", "go", "java"}
	for _, v := range l {
		fmt.Println(v)
	}

	// map型でも問題なくループできる
	sample := map[string]int{"apple": 100, "banana": 200}
	for key, val := range sample {
		fmt.Println(key, val)
	}
	// keyだけ取り出すのであれば、バリューは記述を省略できる
	// 「_」は冗長として認識される
	for key := range sample {
		fmt.Println(key)
	}
	for _, val := range sample {
		fmt.Println(val)
	}
}
