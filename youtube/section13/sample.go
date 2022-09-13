package main

import "fmt"

type Student struct {
	name          string
	math, english float64
}

func main() {
	var s Student
	s.name = "sato"
	s.math = 80
	s.english = 70

	fmt.Println(s)
}
