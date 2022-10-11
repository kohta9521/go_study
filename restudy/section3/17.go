package main

import "fmt"

func main() {
	m := map[string]int{"apple": 100, "banana": 200}
	fmt.Println(m)
	fmt.Println(m["apple"])
	m["banana"] = 3000
	fmt.Println(m)

	m["new"] = 400
	fmt.Println(m)
}
