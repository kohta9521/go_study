package main

import "fmt"

func main() {
	var (
		u8  uint      = 255
		i8  int8      = 127
		f32 float32   = 0, 2
		c64 complex64 = -5 + 1
	)

	fmt.Println(u8, i8, f32, c64)
}
