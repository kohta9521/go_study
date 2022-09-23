package main

import "fmt"

//型
//byte(unit8)型

func main() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	//文字列に変換
	fmt.Println(string(byteA))

	//文字れとをバイトのスライスに変換
	c := []byte("Hi")
	fmt.Println(c)

	//スライスから文字列への変換
	fmt.Println(string(c))
}
