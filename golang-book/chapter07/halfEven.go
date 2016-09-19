package main

import "fmt"

func half(x int) (int, bool) {
	y := x / 2
	return y, y%2 == 0
}

func main() {
	fmt.Println(half(32))
	fmt.Println(half(9))
	fmt.Println(half(150))
}
