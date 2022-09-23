package main

import (
	"fmt"
)

//型
//浮動小数点型

func main() {
	var fl64 float64 = 2.4
	fmt.Println(fl64)

	fl := 3.2
	fmt.Println(fl64 + fl)
	fmt.Printf("%T, %T\n", fl64, fl)

	//基本的には使用しない
	var fl32 float32 = 1.2
	fmt.Printf("%T\n", fl32)

	//型変換
	fmt.Printf("%T\n", float64(fl32))

	//演算が不可能な特殊な型がある 0で割ったりすると無限大になる

	zero := 0.0
	pinf := 1.0 / zero
	fmt.Println(pinf)

	ninf := -1.0 / zero
	fmt.Println(ninf)

	nan := zero / zero
	fmt.Println(nan)
}
