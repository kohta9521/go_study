package main

import "fmt"

func main() {

	x := 10
	y := 12

	if age := x + y; age >= 20 {
		fmt.Println("adult")
	} else if age >= 10 {
		fmt.Println("child")
	} else {
		fmt.Println("kids")
	}

	if x := 12; 5 <= x && x < 10 {
		fmt.Println("child")
	}
}
