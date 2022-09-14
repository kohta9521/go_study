package main

import "fmt"

func main() {

	//数字がキーで値が文字列のマップ
	//HTTPステータスを格納
	hs := map[int]string {
		200: "OK!",
		404: "Not Found",
	}

	//makeで作る
	authors := make(map[string][]string)

	//ブラケットで要素アクセス
	//代入
	authors["Go"] = []string{"Robert Griesemer", "Rob Pile", "Ken Thompson"}

	//データ取得
	status := hs[200]
	fmt.Println(status)
	//"OK"

	//存在しない要素にアクセスするとゼロ値
	fmt.Println(hs[0]) //panic

	//あるかどうかの情報も一緒に取得
	status, ok := hs[304]
	

}