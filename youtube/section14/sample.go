package main

import "fmt"

type Student struct {
	name               string
	math, english float64
}

func (s Student) avg(math, english float64) float64{
	return (math + english) / 2
}

func main() {
	a001 := Student{"sato", 80, 70}
	fmt.Println(a001.avg(80, 70))
}