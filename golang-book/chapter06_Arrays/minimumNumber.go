package main

import "fmt"

func main() {
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	for i := 1; i < len(x); i++ {
		for j := 1; j < len(x); j++ {
			var k int = j - 1
			if x[j] < x[k] {
				jVal := x[j]
				kVal := x[k]
				x[j] = kVal
				x[k] = jVal
			}
		}
	}
	fmt.Println(x)
	fmt.Printf("The minimum number is %d\n", x[0])
}
