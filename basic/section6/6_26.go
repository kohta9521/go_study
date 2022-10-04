package main

import "fmt"

//論理演算子

func main() {
	fmt.Println(true && false == true)
	fmt.Println(true && true == true)

	fmt.Println(true || false == true)

	fmt.Println(!true)
	fmt.Println(!false)
}
