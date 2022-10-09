package main

import (
	"fmt"
)

//main関数の中でしか呼び出せない
func buzz() {
	fmt.Println("Buzz")
}

//main関数よりも先に表示される しばしば初期設定などで使用される
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("Hello World", "TEST")
	buzz()
}
