package main

import "fmt"

type Person struct {
	name string
}

func (p *Person) Talk() {
	fmt.Printf("Hi, my name is %s\n", p.name)
}

type Android struct {
	Person Person
	model  string
}

func (a *Android) phoneTalk() {
	fmt.Printf("I have a %s\n", a.model)
}

type Admin struct {
	Person
	rank string
}

func (admin *Admin) adminSpeak() {
	fmt.Printf("I rank as %s\n", admin.rank)
}

func main() {
	fmt.Println("What's your name?")
	var name string
	fmt.Scanf("%s", &name)
	p := Person{name}
	p.Talk()

	a := Android{Person: p, model: "sony xperia"}
	a.Person.Talk()
	a.phoneTalk()

	admin := Admin{p, "awesome"}
	admin.Talk()
	admin.adminSpeak()
}
