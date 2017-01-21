package main

import (
	greeting "./greeting"
)

func main() {
	var salutationsSlice = greeting.Salutations{
		{"Billy-Bob", "Hello"},
	}
	salutationsSlice.Greet(greeting.CreatePrintFunction("?"), true)

	salutationsSlice[0].Rename("Frog")
	salutationsSlice.Greet(greeting.CreatePrintFunction("?"), true)
}
