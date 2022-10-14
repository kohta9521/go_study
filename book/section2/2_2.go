package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := hello.Iput("type a ptice")
	n, err := strconv.Atoi(x)
	p := float64(n)
	fmt.Println(int(p * 1.1))
}
