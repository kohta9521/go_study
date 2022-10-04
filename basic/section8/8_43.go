package main

import (
	"fmt"
	"time"
)

//go goroutin

func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}
}
