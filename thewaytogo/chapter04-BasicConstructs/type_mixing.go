package main

import "fmt"

func main() {
	var a int
	var b int32
	a = 15
	//b = a + a
	//can't assign int value to int32 variable
	b = b + 5
	fmt.Printf("a is %d\n", a)
	fmt.Printf("b + 5 is %d\n", b)
}
