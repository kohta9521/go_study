package main

import "fmt"

func main() {
	// for i := 0; i <= 7; i++ {
	// 	if i == 6 {
	// 		break
	// 	}
	// 	if i == 3 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }

	arr := [...]int{2, 4, 6, 8, 10}
	sum := 0

	for i := 0; i <= 4; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println(sum)
}
