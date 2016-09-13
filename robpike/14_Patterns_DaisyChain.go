/*
  NOT a loop - just goes all the way around the chain
*/
package main

import (
	"fmt"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	//Launch the constant into the first channel and print it from the last in the chain
	//This runs 100000 goroutines in seconds
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
