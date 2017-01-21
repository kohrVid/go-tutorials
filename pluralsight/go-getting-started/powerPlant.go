package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	//plantCapacities := []float64{30, 30, 30, 60, 60, 100}
	plants := []PowerPlant{
		PowerPlant{hydro, 300, active},
		PowerPlant{wind, 30, active},
		PowerPlant{wind, 25, inactive},
		PowerPlant{wind, 35, active},
		PowerPlant{solar, 45, unavailable},
		PowerPlant{solar, 40, inactive},
	}

	grid := PowerGrid{300, plants}

	if option, err := requestOption(); err == nil {
		fmt.Println("")

		switch option {
		case "1":
			grid.generatePlantReport()
		case "2":
			grid.generatePowerGridReport()
		}
	} else {
		fmt.Println(err.Error())
	}
}
func generatePlantCapacityReport(plantCapacities ...float64) {
	for idx, cap := range plantCapacities {
		fmt.Printf("Plant %d capacity: %.0f\n", idx, cap)
	}
}

func requestOption() (option string, err error) {
	fmt.Println("1) Generate Power Plant Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Print("Please choose and option: ")

	fmt.Scanln(&option)

	if option != "1" && option != "2" {
		err = errors.New("Invalid option selected")
	}
	return
}

type PlantType string

const (
	hydro PlantType = "Hydro"
	wind  PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const (
	active      PlantStatus = "Active"
	inactive    PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity  float64
	status    PlantStatus
}

type PowerGrid struct {
	load   float64
	plants []PowerPlant
}

func (pg *PowerGrid) generatePlantReport() {
	for idx, p := range pg.plants {
		label := fmt.Sprintf("Plant #%d", idx)
		fmt.Println(label)
		fmt.Println(strings.Repeat("-", len(label)))
		fmt.Printf("Type: \t\t%s\n", p.plantType)
		fmt.Printf("Capacity: \t%.0f\n", p.capacity)
		fmt.Printf("Load: \t\t%s\n\n", p.status)
	}
}

func (pg *PowerGrid) generatePowerGridReport() {
	capacity := 0.
	for _, p := range pg.plants {
		if p.status == active {
			capacity += p.capacity
		}
	}
	label := "Power Grid Report"
	fmt.Println(label)
	fmt.Println(strings.Repeat("-", len(label)))
	fmt.Printf("Capacity: \t%.0f\n", capacity)
	fmt.Printf("Load: \t\t%.0f\n", pg.load)
	fmt.Printf("Utilisation: \t%.2f%%\n", pg.load/capacity*100)
}
