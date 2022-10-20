package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}

func main() {
	x := Input("input data")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!!")
}
