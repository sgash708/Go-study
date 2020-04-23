package main

import "fmt"

func aFor() {

	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("continue")
			continue
		}
		if i > 5 {
			fmt.Println("break")
			break
		}
		fmt.Println(i)
	}

	// while文と同じ、初期値はなくてもOK！
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Println(sum)

	// 無限ループ
	for {
		fmt.Println("HELLO!")
	}
}
