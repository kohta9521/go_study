package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	scketches := []string{"Dead Porrot", "Killer joke", "Spanish Inquisition", "Spam"}
	for i, s := range scketches {
		fmt.Println(i, s)
	}

	//1変数だけ書けばインデックスのみを受け取れる
	for i := range scketches {
		fmt.Println(i)
	}

	//ブランク識別子でインデックスを読みとバイsて値だけを扱う
	for _, s := range scketches {
		fmt.Println(s)
	}

	//breakなどで読み飛ばす処理
	for _, s := range scketches {
		if strings.HasPrefix(s, "K") {
			continue
		}
		if strings.HasSuffix(s, "n") {
			break
		}
	}

	//ブール型の要素を１つだけ持つループ文
	counter := 0
	for counter < 10 {
		fmt.Println("ブール値がtrueの間回り続けるループ")
		counter += 1
	}

	end := time.Now().Add(time.Second)
	for {
		fmt.Println("brealやreturnで抜けないと終わらないループ")
		if end.Before(time.Now()) {
			break
		}
	}

	//伝統的なループ処理
	for i := 0; i < 10; i++ {
		fmt.Println("10回繰り替えす")
	}
}
