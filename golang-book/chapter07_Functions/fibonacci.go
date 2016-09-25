package main

import "fmt"

func fib(x int) int {
	if x <= 2 {
		return 1
	} else {
		return (fib(x-1) + fib(x-2))
	}
}

func main() {
	fmt.Println(fib(1))
	fmt.Println(fib(2))
	fmt.Println(fib(3))
	fmt.Println(fib(4))
	fmt.Println(fib(5))
	fmt.Println(fib(6))
	fmt.Println(fib(7))
	fmt.Println(fib(8))
	fmt.Println(fib(9))
	fmt.Println(fib(10))
}
