package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("I shall not print")
	case true:
		fmt.Println("I Will Print")
	default:
		fmt.Println("I print if no other statement is true")
	}
}