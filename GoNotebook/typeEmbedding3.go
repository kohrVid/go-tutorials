package main

import . "fmt"

type Hello struct{}

func (h Hello) String() string {
	return "Hello"
}

type Message struct {
	*Hello
	World string
}

func (v Message) String() (r string) {
	if v.Hello == nil {
		r = v.World
	} else {
		r = Sprintf("%v %v", v.Hello, v.World)
	}
	return
}

func main() {
	m := &Message{}
	Println(m)
	m.Hello = new(Hello)
	Println(m)
	m.World = "world"
	Println(m)
}
