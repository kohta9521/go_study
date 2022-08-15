package main

import "fmt"

func init() {
	fmt.Println("Init")
}

func buzz() {
	fmt.Println("Buzz")
}

func main() {
	buzz()
	fmt.Println("Hello World!")
}
