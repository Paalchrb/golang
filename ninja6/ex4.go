package main

import "fmt"

type person struct{
	first string
	last string
	age int
}

func (p person) speak() {
	fmt.Println("Hello! My name is", p.first, p.last, "and I'm",  p.age,  "years old")
}

func main() {
	p1 := person{
		first: "James",
		last: "Bond",
		age: 34,
	}

	p1.speak()

}