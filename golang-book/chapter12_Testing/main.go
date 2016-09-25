package main

import m "maths"
import "fmt"

func main() {
	xs := []float64{4, 2, 3, 1}
	avg := m.Average(xs)
	fmt.Println(avg)
	fmt.Println(m.Min(xs))
	fmt.Println(m.Max(xs))
}
