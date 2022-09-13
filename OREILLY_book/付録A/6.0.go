package main

import "fmt"

func main() {
	var i int = 10
	var p *int
	p = &i
	fmt.Println(*p)
}