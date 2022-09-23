package main

import (
	"fmt"
)

//整数型

func main() {
	var i int = 100
	fmt.Println(i)

	var i2 int64 = 200

	fmt.Println(i + 50) //150
	// fmt.Println(i + i2) //int型でも64などの部分が違うとエラーが発生する

	//型変換
	//型を表示
	fmt.Printf("%T\n", i2)

	fmt.Printf("%T\n", int32(i2))

	fmt.Println(i + int(i2))

}
