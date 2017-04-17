package main

import . "fmt"

type Hello struct{}

func (h Hello) String() string {
	return "Hello, earth!"
}

type Message struct {
	Hello
}

func main() {
	m := &Message{}
	Println(m.Hello.String())
	Println(m.String())
	Println(m)
}
