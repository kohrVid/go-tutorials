package main

import (
	greeting "./greeting"
	"fmt"
)

func main() {
	salutations := greeting.Salutations{
		{"Billy-Bob", "Hello"},
	}

	fmt.Fprintf(&salutations[0], "The count is %d", 10)

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)
	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}
