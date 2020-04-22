package main

import "fmt"

func aByte() {
	b := []byte{72, 73}
	fmt.Println(b)
	// stringにキャストすると「ASCIIコード」から探し当てて表示する
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
