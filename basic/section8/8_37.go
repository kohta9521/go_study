package main

import (
	"fmt"
)

//for
//繰り返し処理

func main() {
	i := 0
	for {
		i++
		if i == 3 {
			break
		}
		fmt.Println("Loop")
	}

	//条件付きfor

	point := 0
	for point < 10 {
		fmt.Println(point)
		point++
	}

	//古典的for
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		if i == 6 {
			break
		}
		fmt.Println(i)
	}

	//配列でfor文利用
	arr := [3]int{1, 2, 3}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	//範囲式のfor

	for i, v := range arr {
		fmt.Println(i, v)
	}

	//スライス() 配列によく似たもの
	sl := []string{"Python", "PHP", "Go"}
	for i, v := range sl {
		fmt.Println(i, v)
	}

	//mapについて　自社型のようなもの　キーバリュー
	m := map[string]int{"apple": 100, "Banana": 200}

}
