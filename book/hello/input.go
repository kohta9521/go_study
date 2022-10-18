package hello

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

/* package helloが出てきた場合以下のコードを貼り付ける
注意 main関数内のhrllo.input部分を削除する



func Input(msg string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(msg + ": ")
	scanner.Scan()
	return scanner.Text()
}






*/
