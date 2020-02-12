package main

import (
	"fmt"
)

func main() {
	if 1 > 2 {
		fmt.Println("One is bigger than two")
	} else if 3 == 4 {
		fmt.Println("Three equals four")
	} else {
		fmt.Println("No other condition is true")
	}
}