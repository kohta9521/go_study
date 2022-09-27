package main

//関数
//甘雨雨を引数にとる関数
import (
	"fmt"
)

func CallFunction(f func()) {
	f()
}

func main() {
	CallFunction(func() {
		fmt.Println("I'm a function!")
	})
}
