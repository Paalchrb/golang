package main

import "fmt"

func main() {
	favSport := "surfing"
	switch favSport {
	case "soccer":
		fmt.Println("I love soccer")
	case "gymnastics":
		fmt.Println("I love gymnastics")
	case "surfing":
		fmt.Println("I love surfing")
	default:
		fmt.Println("I hate sports") 
	}
}