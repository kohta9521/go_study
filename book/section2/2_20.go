package main

import "fmt"

func main() {
	m := map[string]int{
		"a": 100,
		"b": 200,
		"c": 300,
	}
	m["total"] = m["a"] + m["b"] + m["c"]
	fmt.Println(m)
}
