package main

import "fmt"

func main() {
	l := []string{"Python", "Go", "java"}

	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}

	for _, v := range l {
		fmt.Println(v)
	}

	m := map[string]int{"apple": 100, "banana": 200}

	for k, v := range m {
		fmt.Println(k, v)
	}

	for _, v := range m {
		fmt.Println(v)
	}
}
