package main

import (
	"fmt"
)

var i int = 10

var p *int

p = &i

fmt.Println(*p)