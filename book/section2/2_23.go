package main

import "fmt"

func main() {
	a := []int{10, 20, 30}
	fmt.Println(a)
	a = push(a, 100)
	fmt.Println(a)
	a = pop(a)
	fmt.Println(a)
}

func push(a []int, v int) []int {
	return append(a, v)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}
