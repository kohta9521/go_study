package main

import "fmt"

/*
	この部分は複数行用のコメントアウト
*/

func init() {
	fmt.Println("Init")
}

func buzz() {
	fmt.Println("Buzz")
}

func main() {
	buzz()
	fmt.Println("Hello World!", "TEST TESR")
}
