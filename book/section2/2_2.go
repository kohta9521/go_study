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

// func main() {
// 	x := Iput("type a ptice")
// 	n, err := strconv.Atoi(x)
// 	p := float64(n)
// 	fmt.Println(int(p * 1.1))
// }
