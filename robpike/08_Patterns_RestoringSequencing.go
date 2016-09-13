package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string    //message sent
	wait chan bool //Include a wait chanel
}

func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool) //Shared between all messages

	go func() {
		for i := 0; ; i++ {
			//Block the waitForIt channel
			c <- Message{str: fmt.Sprintf("%s %d", msg, i), wait: waitForIt}
			/*The original Struct syntax didn't work and so had to
			be replaced with the above
			*/
			time.Sleep(time.Duration(rand.Intn(2e3)) * time.Millisecond)
			<-waitForIt
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}

func main() {
	c := fanIn(boring("Joe"), boring("Amy"))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}

}
