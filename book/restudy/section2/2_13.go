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
	t := 0
	x := Input("type a nu,ber")
	n, err := strconv.Atoi(x)
	if err != nil {
		goto err
	}
	for i := 1; i <= n; i++ {
		t += 1
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!!")
}
