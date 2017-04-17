package main

import (
	greeting "./greeting"
	"fmt"
)

func main() {
	salutations := greeting.Salutations{
		{"Billy-Bob", "Hello"},
	}
	fmt.Fprintf(&salutations[0], "The count is %d\n", 10)

	done := make(chan bool)

	go func() {
		salutations.Greet(greeting.CreatePrintFunction("<C>"), true)
		done <- true
	}()

	salutations.Greet(greeting.CreatePrintFunction("?"), true)
	<-done
}
