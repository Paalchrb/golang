package main

import "fmt"

func main() {
	f := foo()
	fmt.Println(f)
	n, s := bar()
	fmt.Println(n, s)
}

func foo() int {
	return 42
}

func bar() (int, string) {
	return 25, "HelloWorld!"
}