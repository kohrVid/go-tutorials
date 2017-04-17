package main

import "fmt"

func main() {
	plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose and option: ")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		for idx, cap := range plantCapacities {
			fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
		}
	case "2":
		capacity := 0.
		for _, plantId := range activePlants {
			capacity += plantCapacities[plantId]
		}
		fmt.Printf("Capacity: \t%.0f\n", capacity)
		fmt.Printf("Load: \t\t%.0f\n", gridLoad)
		fmt.Printf("Utilisation: \t%.1f%%\n", gridLoad/capacity*100)
	default:
		fmt.Println("Unknown option, no action taken")
	}
}
