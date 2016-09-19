package main

import "fmt"

func main() {
	fmt.Print("Enter the temperature in Fahrenheit: ")
	var fahrenheit float64
	fmt.Scanf("%f", &fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("The temperature is %.2f°C\n", celsius)
}
