package main

import (
	"fmt"
	"os/user"
	"time"
)

func main() {
	fmt.Println("Hello Wolrld!", time.Now())
	fmt.Println(user.Current())
}
