package main

import "fmt"

func main() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("Hi")
	fmt.Println(c)
	fmt.Println(string(c))
}
