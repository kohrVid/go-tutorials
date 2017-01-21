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

	done := make(chan bool, 2)

	go func() {
		salutations.Greet(greeting.CreatePrintFunction("<C>"), true)
		done <- true
		done <- true
		fmt.Println("Done")
	}()

	salutations.Greet(greeting.CreatePrintFunction("?"), true)
	<-done
}
