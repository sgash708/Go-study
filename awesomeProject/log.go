package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func loggingOpenFile(LoggingFile string) {
	logfile, _ := os.OpenFile(LoggingFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	// longfile指定すると、フルパスで記述される
	// log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	// shortfile指定すると、相対パスで記述される
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)

}
func aLog() {
	fmt.Println("this is not a log")

	loggingOpenFile("test.log")
	_, err := os.Open("./fafa")
	if err != nil {
		log.Fatalln("there is no such file. exit!!", err)
	}

	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalf("%T %v", "fatal test", "fatal test")
	log.Fatalln("error!")

	// 表示されない
	// 13行目のerror!で処理が終了(exit)しているから
	fmt.Println("ok!!")
}
