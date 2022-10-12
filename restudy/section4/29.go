package main

import (
	"fmt"
	"log"
	"os"
)

func LoggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.OpenFile(logFile, os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiLogFile)
}

func main() {
	//これは他の言語での場合 Goの標準にはない
	/*
		logging.debug("")
		logging.info("")
		logging.warn("")
		logging.error("")
		logging.exception("")
	*/

	// file, err := os.Open("fgafdsage")
	// if err := nil {
	// 	log.Fatalln("Exit", err)
	// }
	log.Println("logging!")
	log.Printf("%T %v", "test", "test")

	log.Fatalln("error!") //この時点でコードが終了する

	fmt.Println("Ok!")
}
