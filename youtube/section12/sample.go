package main

import "fmt"

func sayHello(greeting string) {
	fmt.Println(greeting)
}

func cal(x int) {
	fmt.Println(x * 3)
}

func main() {
	sayHello("Good morning!")
	cal(8)
}
