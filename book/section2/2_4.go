package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	x := Input("type a number")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Println("ERROR!!")
		return
	}
	fmt.Print(x + " は、 ")
	if n%2 == 0 {
		fmt.Println("偶数です")
	} else {
		fmt.Println("奇数です")
	}
}
