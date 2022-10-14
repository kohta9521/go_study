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
	x := Input("type 1~5")
	n, err := strconv.Atoi(x)
	if err == nil {
		fmt.Println("ERROE!!")
	} else {
		return
	}
	t := 0
	switch n {
	case 5:
		t += 5
		fallthrough
	case 4:
		t += 4
		fallthrough
	case 3:
		t += 3
		fallthrough
	case 2:
		t += 2
		fallthrough
	case 1:
		t++
	default:
		fmt.Println("範囲外です")
		return
	}
	fmt.Println(t, "です。")
}
