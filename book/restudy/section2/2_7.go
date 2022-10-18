package main

import (
	"bufio"
	"fmt"
	"os"
)

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	x := 5
	switch x {
	case f(1):
		fmt.Println("* first case. *")
	case f(2):
		fmt.Println("* second case. *")
	case f(3):
		fmt.Println("* third case *")
	default:
		fmt.Println("* default case *")
	}
}

func f(n int) int {
	fmt.Println("No,", n)
	return n
}
