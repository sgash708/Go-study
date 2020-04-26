package main

import (
	"fmt"
)

func thirdPartyConnectDB() {
	// 自身で例外を投げることができる
	// 「ただし」自分でパニックを書くようなコードは記述してはいけない
	// エラーハンドリングを自分できちんと行うことが大事
	panic("Unable to connect database!!")
}

func save() {
	defer func() {
		// 「recover」でシステムを強制終了させないようにする
		s := recover()
		fmt.Println(s)
	}()
	// 「defer」の処理も前に下記記述を行うと「recover」されない
	thirdPartyConnectDB()
	// 本来であれば、「thirdPartyConnectDB()」に返り値を持たせること
	// if err := thirdPartyConnectDB(); err != nil {
	// 	log.Fatalln("Error DB")
	// }
}

func aPanicRecover() {
	save()
	fmt.Println("OK?")
}
