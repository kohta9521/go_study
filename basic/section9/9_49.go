package main

import (
	"fmt"
)

//スライス
//for

func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	//インデックス番号を使用しない場合
	// for _, v := range sl {
	// 	fmt.Println(v)
	// }

	//古典的forを用いた場合
	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}
}
