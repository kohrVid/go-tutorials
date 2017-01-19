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
		name := strings.ToUpper(os.Args[1])
		suffix := dedupe(name[1:])
		scoreString := handleCharacters(suffix)
		scoreSlice := strings.Split(scoreString, ",")
		fmt.Printf("Your score is: %v%v\n", name[:1], strings.Join(handleScore(scoreSlice), ""))

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
		"A", "",
		"E", "",
		"I", "",
		"O", "",
		"U", "",
		"Y", "",
		"H", "",
		"W", "",
		"B", "1,",
		"F", "1,",
		"P", "1,",
		"V", "1,",
		"C", "2,",
		"G", "2,",
		"J", "2,",
		"K", "2,",
		"Q", "2,",
		"S", "2,",
		"X", "2,",
		"Z", "2,",
		"D", "3,",
		"T", "3,",
		"L", "4,",
		"M", "5,",
		"N", "5,",
		"R", "6,",
	)
	return replacer.Replace(input)
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
