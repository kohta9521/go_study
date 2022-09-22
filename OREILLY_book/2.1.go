package main

import (
	"fmt"
)

// x := 1

if condition {
	//コードのブロックの内側で変数,x を再宣言できる。
	x := 2
	fmt.Println("x = ", x)
}

// condition がtrueの場合でも、ここの位置では必ず１になる。
fmt.Println("x = ", x)