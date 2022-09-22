package main

import "fmt"

func main() {
	var s string = "running"

	switch s {
	case "running":
		fmt.Println("実行中")
	case "stop":
		fmt.Println("停止中")
	default:
		fmt.Println("その他")
	}
}
