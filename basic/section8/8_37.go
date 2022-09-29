package main

import "fmt"

//for
//繰り返し構文

func main() {
	// i := 0
	// for {
	// 	i++
	// 	if i == 3 {
	// 		break
	// 	}
	// 	fmt.Println("Loop")
	// }

	// point := 0
	// for point < 10 {
	// 	fmt.Println(point)
	// 	point++
	// }

	// for i := 0; i < 10; i++ {
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	if i == 6 {
	// 		break
	// 	}
	// 	fmt.Println(i)
	// }

	// arr := (3)int{1, 2, 3}
	// for i := 0; i < len(arr); u L++

	// arr := [3]int{1, 2, 3}

	// for i, v := range arr {
	// 	fmt.Println(i, v)
	// }

	// sl := []string{"Python", "PHP", "Go"}
	// for i, v := range sl {
	// 	fmt.Println(i, v)
	// }

	m := map[string]int{"apple": 100, "Banana": 200}

	for k := range m {
		fmt.Println(k)
	}

}
