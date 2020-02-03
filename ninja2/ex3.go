package main

import "fmt"

const (
	a = 31
	b int = 41
)

func main() {
	fmt.Printf("%d\t%T\n", a, a)
	fmt.Printf("%d\t%T\n", b, b)
}