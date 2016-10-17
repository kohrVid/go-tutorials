package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("One for loop:")
	for i := 1; i <= 25; i++ {
		line := strings.Repeat("G", i)
		fmt.Printf("%s\n", line)
	}
}
