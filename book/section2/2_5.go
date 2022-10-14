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

// func man() {
// 	x := Input("type a number")
// 	fmt.Print(x + " は、 ")
// 	if n, err := strconv.Atoi(x); err != nil {
// 		fmt.Println("ERROR!!")
// 		return
// 	}
// 	if n%2 == 0 {
// 		fmt.Println("偶数です")
// 	} else {
// 		fmt.Println("奇数です")
// 	}
// }
