package main

import "fmt"

type Vertex struct {
	X, Y int
	// X int
	// // 小文字だと「private」になる
	// Y int

	S string
}

func chVertex(v Vertex) {
	v.X = 1000
}
func chVertex2(v *Vertex) {
	v.X = 1000
	// 本来はこの形でかくけど、構造体は認識してくれるので「(*v)」とかかない
	// (*v).X = 1000
}

func aStruct() {
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)        // {1 2}
	fmt.Println(v.X, v.Y) // オブジェクト思考的に使用可能
	v.X = 100
	fmt.Println(v, v.X, v.Y)

	// 片方しか入れなかったらどうなるか
	v2 := Vertex{X: 1}
	fmt.Println(v2) // {1 0}

	// 変数省略の場合
	v3 := Vertex{1, 2, "test"}
	fmt.Println(v3) // {1 2 test}

	v4 := Vertex{}
	// fmt.Println(v4) // {0 0 }
	fmt.Printf("%T %v\n", v4, v4)

	// Vertexの宣言は「nill」にならない
	var v5 Vertex
	// fmt.Println(v5)
	fmt.Printf("%T %v\n", v5, v5)

	v6 := new(Vertex)
	// fmt.Println(v6)
	fmt.Printf("%T %v\n", v6, v6)

	// newと同じ動きでポインタが帰ってくる
	// 一目見て「アドレス」が返ってくることがわかるので「おすすめ」
	v7 := &Vertex{}
	fmt.Printf("%T %v\n", v7, v7)

	// mapやsliceに関しては「make」「new」を使う
	// vertexに関しては、どっちでも良い

	vex := Vertex{1, 2, "test"}
	chVertex(vex)
	fmt.Println(vex)

	// 書き換えたい時には、アドレスをつけて宣言する
	vex2 := &Vertex{1, 2, "test"}
	chVertex2(vex2)
	fmt.Println(vex2)
	// 実態をみるときは「*」をつける
	fmt.Println(*vex2)

}
