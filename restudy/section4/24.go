package main

import (
	"fmt"
)

func by2(num int) string {
	if num%2 == 0 {
		return "OK"
	} else {
		return "NO"
	}
}

func main() {
	result := by2(10)

	if result == "OK" {
		fmt.Println("great")
	}

	if result2 := by2(10); result2 == "OK" {
		fmt.Println("great 2")
	}

	/*
		num := 6
		if num%2 == 0 {
			fmt.Println("By 2")
		} else if num%3 == 0 {
			fmt.Println("By 3")
		} else {
			fmt.Println("else")
		}
	*/

	x, y := 10, 10
	if x == 10 && y == 10 {
		fmt.Println("%%")
	}
	if x == 10 || y == 10 {
		fmt.Println("||")
	}
}
