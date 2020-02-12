package main

import "fmt"

func main() {
	person := struct{
		first string
		last string
		age int
		hobbies []string
	}{
		first: "Peter",
		last: "Pan",
		age: 11,
		hobbies: []string{
			"Fencing",
			"Flying",
			"Bully captain hook",
		},
	}

	fmt.Println(person);
}