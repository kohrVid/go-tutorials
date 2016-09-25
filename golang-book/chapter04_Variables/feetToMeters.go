package main

import "fmt"

func main() {
	fmt.Print("Enter the length in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)

	meters := feet * 0.3048
	fmt.Printf("The length in meters is %.2f m\n", meters)
}
