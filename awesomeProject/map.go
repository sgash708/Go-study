package main

import "fmt"

func aMap() {
	// mapはC#のdictionaryと似ている
	// ex) map[keyの型]valueの型
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 300
	fmt.Println(m)

	// 新しい要素追加
	m["new"] = 500
	fmt.Println(m)

	// 存在しないキーを指定した場合
	fmt.Println(m["nothing"]) // 0と表示

	// 指定したキーが存在しているかチェックする
	// 「ok」は存在をチェックするため
	// 上記の変数名はなんでも良い
	v, ok := m["apple"]
	fmt.Println(v, ok)

	// 存在しないキーでチェック
	v2, ok2 := m["nothing"]
	fmt.Println(v2, ok2)

	m2 := make(map[string]int)
	m2["pc"] = 5000
	fmt.Println(m2)

	// 宣言だけだと「Nil」になる
	var sss []int
	if sss == nil {
		fmt.Println("Nil")
	}
}
