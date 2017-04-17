package main

import . "fmt"

type Message struct {
	X string
	y *string
}

func (v Message) Print() {
	if v.y != nil {
		Println(v.X, *v.y)
	} else {
		Println(v.X)
	}
}

func (v *Message) Store(x, y string) {
	v.X = x
	v.y = &y
}

func main() {
	m := &Message{}
	m.Print()
	m.Store("Hello", "earthling")
	m.Print()
}
