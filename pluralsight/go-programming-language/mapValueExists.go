package main

import "fmt"

func main() {
	fmt.Println(GetPrefix("Jo"))
}

func GetPrefix(name string) (prefix string) {
	prefixMap := map[string]string{
		"Bob":      "Mr. ",
		"Jo":       "Dr. ",
		"Arethusa": "Ms. ",
	}

	if value, exists := prefixMap[name]; exists {
		fmt.Println(exists)
		return value
	}
	return "..."
}
