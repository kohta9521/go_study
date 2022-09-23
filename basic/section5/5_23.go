package main

import "fmt"

//定数

//グローバルに記述
const Pi = 3.14

const (
	URL      = "http://xxx.com"
	SiteName = "test"
)

//連続する整数の演算を生成
const (
	c0 = iota
	c1
	c2
)

const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)
	fmt.Println(URL, SiteName)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(c0, c1, c2)
}
