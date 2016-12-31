package main

import . "fmt"

type Hello bool

func (h Hello) string() (r string) {
	if h {
		r = "Hello, there!"
	}
	return
}

type Message struct {
	Hello
}

func main() {
	m := &Message{Hello: true}
	Println(m)
	m.Hello = false
	Println(m)
	m.Hello = true
	Println(m)
}
