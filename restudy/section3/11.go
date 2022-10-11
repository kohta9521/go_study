package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello world")
	fmt.Println("Hello " + "World")
	// これはアスキーコード表記
	fmt.Println("Hello World"[0])
	// stringに変換
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "X", 1))
	//実際は置き換わってない
	fmt.Println(s)
	s = strings.Replace(s, "H", "X", 1)
	fmt.Println(s)

	//Contains 文字列が含まれているか
	fmt.Println(strings.Contains(s, "World"))

	//文字列のリテラル
	fmt.Println("Test\n" +
		"Test")

	fmt.Println("\"")
	fmt.Println(`"`)
}
