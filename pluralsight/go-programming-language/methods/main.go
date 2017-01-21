package main

import (
	greeting "./greeting"
)

func main() {
	var salutationsSlice = greeting.Salutations{
		{"Billy-Bob", "Hello"},
	}
	salutationsSlice.Greet(greeting.CreatePrintFunction("?"), true)
}
