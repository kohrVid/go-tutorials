package main

import (
	greeting "./greeting"
)

func main() {
	var s = greeting.Salutation{"Billy-Bob", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("?"), true)
}
