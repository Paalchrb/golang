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

	fmt.Println(p1.first);
	fmt.Println(p1.last);
	for i, v := range p1.favIceCream {
		fmt.Println(i, v)
	}
	
	fmt.Println(p2.first);
	fmt.Println(p2.last);
	for i, v := range p2.favIceCream {
		fmt.Println(i, v)
	}

}