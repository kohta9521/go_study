package main

import "fmt"

//関数
//無名関数

func main() {
	f := func(x, y int) int {
		return x + y
	}
	i := f(1, 2)
	fmt.Println(i)

	//簡易的に記述
	i2 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i2)

}
