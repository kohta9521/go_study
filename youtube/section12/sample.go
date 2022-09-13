package main

import "fmt"

func sayHello(greeting string) {
	fmt.Println(greeting)
}

func cal(x, y int) int {
	return (x / y)
}

func main() {
	sayHello("Good morning!")
	cal(8, 3)
}
