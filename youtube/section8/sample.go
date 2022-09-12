package main

import "fmt"

func main() {
	a := [...]string{"sato", "suzuki", "takahashi"}
	a[1] = "tanaka"

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
}
