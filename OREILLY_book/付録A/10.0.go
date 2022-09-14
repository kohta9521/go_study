package main

import "fmt"

func main() {
	if statusCode == 200 {
		fmt.Println("no error")
	} else if statusCode < 500 {
		fmt.Println("client error")
	} else {
		fmt.Println("server error")
	}


	if result, ok := cache[input]; ok {
		fmt.Println("cached value", result)
	}


}