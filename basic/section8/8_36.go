package main

//if
//条件分岐
//エラーハンドリング
import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error")
	}
	fmt.Printf("i = %d\n", i)

}
