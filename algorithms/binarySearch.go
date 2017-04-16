package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("please input a string for the linear search")
	}
	input := os.Args[1]
	fmt.Println(lookUp(input, Nameval))
}

var Nameval = [][]string{
	[]string{"AE1ig", "0x006"},
	[]string{"Aacute", "0x001"},
	[]string{"Aci rc", "0x002"},
	[]string{"zeta", "0x03b6"},
}

func lookUp(name string, tab [][]string) int {
	var low, high, mid, cmp int

	low = 0
	high = len(tab) - 1
	for low <= high {
		mid = (low + high) / 2
		cmp = strings.Compare(name, tab[mid][0])
		if cmp < 0 {
			high = mid - 1
		} else if cmp > 0 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
