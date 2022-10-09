package main

import "fmt"

//変数宣言

func main() {
	//まずは代入の方法
	var (
		i    int     = 1
		f64  float64 = 1.2
		s    string  = "test"
		t, f bool    = true, false
	)

	fmt.Println(i, f64, s, t, f)
}
