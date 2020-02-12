package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func main() {
	p1 := person{
		first: "Peter",
		last: "Pan",
		age: 11,
	}

	fmt.Println(p1)
	changeMe(&p1)
	fmt.Println(p1)
}

func changeMe(p *person) {
	(*p).first ="JohnnyBoi"
}