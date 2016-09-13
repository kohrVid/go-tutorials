/*
  Handle multiple channels like a switch where each case is a communication
  Blocks until one communication can proceed
  Chooses psuedo randomly
  The default clause rexecutes immediately is no channels are ready
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for { //Only one for block needed
			select {
			//If there are multiple channels available at once, one is selected pseudo-randomly
			case s := <-input1:
				c <- s
				//fmt.Printf("received %v from c1\n", s)
			case s := <-input2:
				c <- s
				//fmt.Printf("received %v from c2\n", s)
			default:
				fmt.Printf("no one was ready to communicate\n")
			}
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Amy"))
	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You're both boring. I'm leaving!")
}
