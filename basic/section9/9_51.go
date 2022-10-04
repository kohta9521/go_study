package main

import (
	"fmt"
)

//マップ

func main() {
	//Pythonでいう辞書型
	var m = map[string]int{"A": 100, "B": 200}

	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}

	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	//make関数を使用してmap作成
	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "JAPAN"
	m4[2] = "USA"

	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])

	//登録されてないキー　から文字変換
	fmt.Println(m4[3])

	s, ok := m4[3]
	if !ok {
		fmt.Println(s, ok)
	}
	fmt.Println(s)

	//map
	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))
}
