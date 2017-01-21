package main

import "fmt"

type Salutation struct {
	name     string
	greeting string
}

func Greet(salutation Salutation, do func(string)) {
	message, alternate := CreateMessage(salutation.name, salutation.greeting, "yo")
	do(message)
	do(alternate)
}

func CreateMessage(name string, greeting ...string) (message string, alternate string) {
	message = greeting[1] + " " + name
	alternate = "HEY!!!!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func main() {
	var s = Salutation{"Billy-Bob", "Hello"}
	Greet(s, PrintLine)
}
