package main

import (
	"fmt"
	"strconv"
)

//型変換

func main() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fll64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	//文字列から数字
	var s string = "100"
	fmt.Printf("s = %T\n", s)

	//数値の１００に変換したい
	i, _ := strconv.Atoi(s)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(i)
	fmt.Printf("i = %T\n", i)

	//数値型から文字列型
	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

	//文字列からバイト配列
	var h string = "Hello World"
	b := []byte(h)
	fmt.Println(b)

	h2 := string(b)
	fmt.Println(h2)

}
