package main

import "fmt"

func main() {
	hs := map[int]string{
		200: "OK",
		404: "Not Found",
	}

	//makeで作る
	authors := make(map[string][]string)

	//ブラケットで要素アクセス
	authors["Go"] = []string{"Rpbert Griesemer", "Rob Pike", "Ken Thompson"}

	//データ取得
	status := hs[200]
	fmt.Println(hs[0])

	status, ok := hs[304]
	status = ""
	ok = false
}
