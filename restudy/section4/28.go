package main

import (
	"fmt"
	"os"
)

func foo() {
	defer fmt.Println("World foo")

	fmt.Println("Hello foo")
}

func main() {

	// defer fmt.Println("World")
	// foo()

	// fmt.Println("Hello")

	//スタッキングでファー　deferの中でも後ろから実行

	/*
		fmt.Println("run")
		defer fmt.Println(1)
		defer fmt.Println(2)
		defer fmt.Println(3)
		fmt.Println("success")
	*/

	file, _ := os.Open("./27.go")
	defer file.Close()
	data := make([]byte, 100)
	file.Read(data)
	fmt.Println(string(data))
}
