package main

import "fmt"

func main() {
	// var t, f bool = true, false
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, t, 1)
	fmt.Printf("%T %v %t\n", f, f, 0)

	fmt.Println(true && true)
	fmt.Println(false && true)
	fmt.Println(false && false)

	fmt.Println(true || true)
	fmt.Println(false || true)
	fmt.Println(false || false)

	fmt.Println(!true)
	fmt.Println(!false)
}
