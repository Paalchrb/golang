package main

import "fmt"

func main() {
	pplMap := map[string][]string{
		"bond_james": []string{"Shaken, not stirred", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr": []string{"Being evil", "Ice cream", "Sunsets"},
	}

	pplMap["octopussy"] = []string{"Killing James", "Killing James", "Make love to James", "Killing James"}

	for key, arr := range pplMap {
		fmt.Println("This is the record for: ", key)
		for i, str := range arr {
			fmt.Printf("\tIndex position: %v\t Value: %v\n", i, str)
		}
	}
}