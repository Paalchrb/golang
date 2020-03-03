package main

import "fmt"

func main() {
	fmt.Println("11 - 5 =", subFive(11))
	fmt.Println("23 - 5 =", subFive(23))
	fmt.Println("35 - 5 =", subFive(35))
}

func subFive(n int) int {
	return n - 5
}