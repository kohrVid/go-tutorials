package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		usageMessage()
	} else {
		name := os.Args[1]
		suffix := name[1:]
		scoreString := handleCharacters(suffix)
		scoreSlice := strings.Split(scoreString, ",")

		fmt.Printf("Your score is: %v%v\n", strings.ToUpper(name[:1]), strings.Join(handleScore(scoreSlice), ""))

	}
}

func usageMessage() {
	fmt.Printf("usage: %s <name>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func handleCharacters(input string) string {
	replacer := strings.NewReplacer(
		"a", "",
		"e", "",
		"i", "",
		"o", "",
		"u", "",
		"y", "",
		"h", "",
		"w", "",
		"b", "1,",
		"f", "1,",
		"p", "1,",
		"v", "1,",
		"c", "2,",
		"g", "2,",
		"j", "2,",
		"k", "2,",
		"q", "2,",
		"s", "2,",
		"x", "2,",
		"z", "2,",
		"d", "3,",
		"t", "3,",
		"l", "4,",
		"m", "5,",
		"n", "5,",
		"r", "6,",
	)
	return replacer.Replace(input)
}

func handleScore(arr []string) []string {
	var three []string
	if len(arr) >= 3 {
		three = arr[0:3:3]
	} else {
		for i := range three {
			three[i] = arr[i]
		}
	}
	for i := range three {
		if three[i] == "" {
			three[i] = "0"
		}
	}
	fmt.Println(len(arr))
	return three
}
