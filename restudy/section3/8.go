package main

import "fmt"

//変数宣言
//まずは代入の方法
var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func main() {
	//まずは代入の方法
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)

	fmt.Println(i, f64, s, t, f)

	foo()
}

//適当な関数を設定してみた
func foo() {
	//関数内部でのみ設定可能
	xi := 1
	xf64 := 1.2
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf64, xs, xt, xf)

	fmt.Printf("%T", xf64)
}
