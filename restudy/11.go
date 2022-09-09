package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello" + " World")
	fmt.Println(string("Hello World"[0]))

	fmt.Println(string.Replace(s, "H", "X", 1))
}
