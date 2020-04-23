package main

import (
	"fmt"
	"time"
)

func getOsName() string {
	return "dadadadadada"
}
func aSw() {
	// 二回以上該当する変数を使う時
	// os := getOsName()
	// switch os {

	// 一度しか使わない変数の場合
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("mac")
		// 「break」文は多言語と違い不要
		// break
	case "windows":
		fmt.Println("windows")
	default:
		fmt.Println("other", os)
	}

	// switchの条件文がなくても成立するものがある
	// Struct(構造体)は条件式に含めることができない
	ti := time.Now()
	switch {
	case ti.Hour() < 12:
		fmt.Println("good mornin!")
	default:
		fmt.Println("good evening")
	}
}
