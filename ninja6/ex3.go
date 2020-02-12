package main

import "fmt"

func main() {
	foo()
	defer foo1()
	foo2()
}

func foo() {
	fmt.Println("Hello1")
}

func foo1() {
	fmt.Println("Hello2")
}

func foo2() {
	fmt.Println("Hello3")
}