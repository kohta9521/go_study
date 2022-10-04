package main

import (
	"fmt"
)

//cannel
//複数のゴールチン間でのデータ受け渡しをするために設計されたデータ構造
//宣言、操作

func main() {
	//チャネルの宣言（双方向可能）
	var ch1 chan int

	//受信専用
	// var ch2 <-chan int

	//送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)

	ch2 := make(chan int)

	//サイズの確認
	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1 //チャネル３に一を送信

	//チャネルの要素数を確認
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("lem", len(ch3))

	//チャネルからデータを受信
	i := <-ch3
	fmt.Println(i)
	fmt.Println("lem", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("lem", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("lem", len(ch3))
}
