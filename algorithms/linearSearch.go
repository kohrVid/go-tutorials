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
	fmt.Println(lookUp(input, flab))
}

func lookUp(word string, wordList []string) int {
	for i := range wordList {
		if strings.Compare(wordList[i], word) == 0 {
			return i
		}
	}
	return -1
}

var flab = []string{
	"actually",
	"just",
	"quite",
	"really",
}
