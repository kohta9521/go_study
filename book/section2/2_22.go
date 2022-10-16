package main

func main() {

}

func push(a []int, v int) []int {
	return append(a, v)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}
