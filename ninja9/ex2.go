package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p *person) speak() {
	fmt.Println("Hello. I am", p.first, p.last, "and I am ", p.age, "years old!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{
		first: "Peter",
		last:  "Pan",
		age:   12,
	}

	saySomething(&p1)

	//Wont run if uncommented
	//saySomething(p1)
}
