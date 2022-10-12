package main

import (
	"fmt"
	"time"
)

//switch文

func getOsName() string {
	return "macasdf"
}

func main() {
	switch os := getOsName(); os {
	case "mac":
		fmt.Println("Mac!")
	case "windows":
		fmt.Println("Windows!")
	default:
		fmt.Println("Default!", os)
	}

	// fmt.Println(os)

	//オブジェクト t
	t := time.Now()
	fmt.Println(t.Hour())

	//switch内に書かないでいい方法
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("afternoon")
	}

}
