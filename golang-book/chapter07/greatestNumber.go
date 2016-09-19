package main

import "fmt"

func greatestNumber(args ...int) int {
	for i := 1; i < len(args); i++ {
		for j := 1; j < len(args); j++ {
			var k int = j - 1
			if args[j] > args[k] {
				jVal := args[j]
				kVal := args[k]
				args[j] = kVal
				args[k] = jVal
			}
		}
	}
	return args[0]

}

func main() {
	fmt.Println(greatestNumber(9, 0, 1))
	fmt.Println(greatestNumber(9, 43, 105, 32))
}
