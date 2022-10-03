package main

import (
	"fmt"
)

//switch
//式スイッチ

func main() {
	n := 1
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("I don't know")
	}

	a := 6
	switch {
	case a > 0 && a < 4:
		fmt.Println("0 < a < 4")
	case a > 3 && a < 7:
		fmt.Println("3 < a < 7")
	}

	switch n := 2; n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	case n < 3 && n < 6:
		fmt.Println("3 > n < 6")
	default:
		fmt.Println("I don't know")
	}
}
