package main

import "fmt"

func main() {
	var plantCapacities []float64

	plantCapacities = []float64{30, 30, 30, 60, 60, 100}

	var capacity float64 = plantCapacities[0] + plantCapacities[1] +
		plantCapacities[2] + plantCapacities[3] + plantCapacities[4] +
		plantCapacities[5]

	var gridLoad = 75.

	utilisation := gridLoad / capacity
	fmt.Printf("Capacity: \t%.0f\n", capacity)
	fmt.Printf("Load: \t\t%.0f\n", gridLoad)
	fmt.Printf("Utilisation: \t%.1f%%\n", utilisation)
}
