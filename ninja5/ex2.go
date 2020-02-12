package main

import "fmt"

type person struct{
	first string
	last string
	favIceCream []string
}

func main() {
	p1 := person {
		first: "Peter",
		last: "Pan",
		favIceCream: []string{
			"strawberry", 
			"banana", 
			"coke",
		},
	}

	p2 := person {
		first: "Mary",
		last: "Madson",
		favIceCream: []string{
			"chocolate", 
			"pistach", 
			"heroin",
		},
	}

	m := map[string] person {
		p1.last: p1,
		p2.last: p2,
	}

	for _, p := range m {
		fmt.Println(p.first)
		fmt.Println(p.last)
		for i, v := range p.favIceCream {
			fmt.Println(i, v)
		}
		fmt.Println(("------------"))
	}
}