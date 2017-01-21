package greeting

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation Salutation, do Printer, isFormal bool) {
	message, alternate := CreateMessage(salutation.Name, salutation.Greeting)

	if prefix := GetPrefix(salutation.Name); isFormal {
		do(prefix + message)
	}
	do(alternate)
}

func GetPrefix(name string) (prefix string) {
	switch {
	case name == "Billy-Bob":
		prefix = "Mr. "
	case name == "Jo", name == "Ulrich", len(name) == 10:
		prefix = "Dr. "
	case name == "Jess":
		prefix = "Ms. "
	default:
		prefix = ""
	}
	return
}

func CreateMessage(name string, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!!!!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
