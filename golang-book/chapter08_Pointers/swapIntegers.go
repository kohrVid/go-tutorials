package main

import "fmt"

func swap(x, y *int) *int {
	xOld := *x
	*x = *y
	*y = xOld
	return x
}

func main() {
	x := 1
	y := 2
	fmt.Println("Original values:")
	fmt.Printf("x=%d\n", x)
	fmt.Printf("y=%d\n", y)
	swap(&x, &y)
	fmt.Println("Values after a swap:")
	fmt.Printf("x=%d\n", x)
	fmt.Printf("y=%d\n", y)
}
