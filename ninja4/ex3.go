package main

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	a := x[:5]
	fmt.Println(a)
	a = x[5:]
	fmt.Println(a)
	a = x[2:7]
	fmt.Println(a)
	a = x[1:6]
	fmt.Println(a)
}