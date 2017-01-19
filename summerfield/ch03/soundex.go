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
		suffix := dedupe(name[1:])
		scoreString := handleCharacters(suffix)
		scoreSlice := strings.Split(scoreString, ",")
		fmt.Printf("Your score is: %v%v\n", strings.ToUpper(name[:1]), strings.Join(handleScore(scoreSlice), ""))

	}
}

func usageMessage() {
	fmt.Printf("usage: %s <name>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

func dedupe(name string) string {
	nameArray := strings.Split(name, "")
	for i := 1; i < len(nameArray); i++ {
		j := i - 1
		if nameArray[i] == nameArray[j] {
			nameArray[i] = ""
		}
	}
	return strings.Join(nameArray, "")
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
	return replacer.Replace(strings.ToLower(input))
}

func handleScore(arr []string) []string {
	three := make([]string, 3, 3)
	var iter []string
	if len(arr) > len(three) {
		iter = three
	} else {
		iter = arr
	}
	for i := range iter {
		three[i] = arr[i]
	}
	if sliceContains("", arr) {
		for i := 0; i < 3; i++ {
			if three[i] == "" {
				three[i] = "0"
			}
		}
	}
	return three
}

func sliceContains(item string, slice []string) bool {
	for _, i := range slice {
		if i == item {
			return true
		}
	}
	return false
}
