package main

import "fmt"

func main() {
	var {
		ErroTooLong          = errors.New("bufio.Scanner: token too long")
		ErrNegativeAdvance  = errors.New("bufio.Scanner: SplitFanc returns negative advance count")
		ErrAdvanceTooFar      = errors.New("bufio.Scanner: SplitFunc returns advance count beyond
		input")
	}
}