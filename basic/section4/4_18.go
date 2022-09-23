package main

import "fmt"

//配列型

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B", "C"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	//要素すうの省略
	arr4 := [...]string{"c", "d"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	//値の取り出し
	fmt.Println(arr2[0])

	var arr5 [4]int
	fmt.Println(arr5)

	fmt.Println(len(arr3))
}
