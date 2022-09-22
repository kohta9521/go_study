package main

import (
	"fmt"
)

func main() {

	var statusCode int = 200

	if statusCode == 200 {
		fmt.Println("no error")
	} else if statusCode < 500 {
		fmt.Println("client erro")
	} else {
		fmt.Println("erverError")
	}

	// if result, ok := cache[input]; ok {
	// 	for i, s := range scketches {
	// 		fmt.Println(i, s)
	// 	}
	// }

}
