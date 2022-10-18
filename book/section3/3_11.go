package main

import (
	"fmt"
)

// Mydata is structure
type Mydata struct {
	Name string
	Dara []int
}

// / PrintData is println all data
func (md Mydata) PrintData() {
	fmt.Println("*** mydata")
}
