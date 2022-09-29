package main

import (
	"fmt"
)

//switch
//型スイッチ

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	anything("aaa")
	anything(1)

	//型アサーション

	var x interface{} = 3
	i := x.(int)
	fmt.Println(i + 2)
	//fmt.Println(x + 2)

	f := x.(float64)
	fmt.Println(f)

	//停止しない方法
	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, "isString")
	} else {
		fmt.Println("I don't know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't know")
	}
}
