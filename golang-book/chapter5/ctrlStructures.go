package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Printf("%d - ", i)
		if i%2 == 0 {
			fmt.Printf("even\n")
		} else {
			fmt.Printf("odd\n")
		}
		i += 1
	}
}
