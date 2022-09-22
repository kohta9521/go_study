package main

import "fmt"

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
}
