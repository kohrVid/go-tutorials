/*
  time.After returns a channels that blocks for the specified duration
  After the interval, the chnnels delivers the current time once.
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
			case s := <-input2:
				c <- s
			default:
				fmt.Printf("no one was ready to communicate\n")
			}
		}
	}()
	return c
}

func main() {
	c := boring("Joe")
	timeout := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You're too slow and you talk too damned much.")
			return
		}
	}
}
