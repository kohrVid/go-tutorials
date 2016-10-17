package main

import "fmt"

func main() {
	N := 9
	for i, j := 0, N; i < j; i, j = i+1, j-1 {
		fmt.Printf("j is %d and i is %d\n", j, i)
	}
}
