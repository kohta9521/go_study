package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("hello world!")
	fmt.Println("hello" + "world")
	fmt.Println(string("Hello World"[0]))

	var s string = "Hello World"
	fmt.Println(strings.Replace(s, "H", "H", 1))

	fmt.Println(strings.Contains(s, "World"))

	fmt.Println("Test\n" +
		"Test")

	fmt.Println("\"")
	fmt.Println(`\`)
}
