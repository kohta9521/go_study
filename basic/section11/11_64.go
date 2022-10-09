package main

import "fmt"

type T struct {
	User User
}

type User struct {
	Name string
	Age  int
	//X, Y int
}

//構造体の埋め込み

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)
	fmt.Println(t.User.Name)
}
