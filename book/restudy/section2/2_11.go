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
	if err == nil {
		fmt.Print("1から" + x + "の合計は、")
	} else {
		return
	}
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println(t, "です。")
}
