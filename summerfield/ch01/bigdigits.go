package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if os.Args[1] == "--help" {
		helpMessage()
	} else {
		var stringOfDigits string //:= stringArgs(os.Args)
		switch len(os.Args) {
		case 1:
			helpMessage()
		case 2:
			stringOfDigits = os.Args[1]
			generateDigit(stringOfDigits)
		case 3:
			if os.Args[1] == "--bar" || os.Args[1] == "-b" {
				stringOfDigits = os.Args[2]
				bar(stringOfDigits)
			} else {
				fmt.Println("Sorry don't know that one!")
				helpMessage()
			}

		}
	}
}

func bar(stringOfDigits string) {
	barLength := 6 * len(stringOfDigits)
	fmt.Printf(strings.Repeat("*", barLength))
	fmt.Printf("**\n")
	generateDigit(stringOfDigits)
	fmt.Printf("\n")
	fmt.Printf(strings.Repeat("*", barLength))
	fmt.Printf("**\n")
}

func generateDigit(stringOfDigits string) {
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			digit := stringOfDigits[column] - '0'
			if 0 <= digit && digit <= 9 {
				line += bigDigits[digit][row] + " "
			} else {
				log.Fatal("invalid whole number")
			}
		}
		fmt.Println(line)
	}
}

func helpMessage() {
	fmt.Printf("usage: %s[-b|--bar] <whole-number>\n", filepath.Base(os.Args[0]))
	os.Exit(1)
}

var bigDigits = [][]string{
	{
		"  000  ",
		" 0   0 ",
		"0     0",
		"0     0",
		"0     0",
		" 0   0 ",
		"  000  ",
	},
	{" 1 ", "11 ", " 1 ", " 1 ", " 1 ", " 1 ", "111"},
	{" 222 ", "2   2", "   2 ", "  2  ", " 2   ", "2    ", "22222"},
	{" 333 ", "3   3", "    3", "  33 ", "    3", "3   3", " 333 "},
	{"   4  ", "  44  ", " 4 4  ", "4  4  ", "444444", "   4  ", "   4  "},
	{"55555", "5    ", "5    ", " 555 ", "    5", "5   5", " 555 "},
	{" 666 ", "6    ", "6    ", "6666 ", "6   6", "6   6", " 666 "},
	{"77777", "    7", "   7 ", "  7  ", " 7   ", "7    ", "7    "},
	{" 888 ", "8   8", "8   8", " 888 ", "8   8", "8   8", " 888 "},
	{" 9999", "9   9", "9   9", " 9999", "    9", "    9", "    9"},
}
