package main

import "fmt"

func main() {
	greeting := func(name string) string {
		return fmt.Sprint("Hello there, ", name);
	}("Peter")

	fmt.Println(greeting);

	fgreet := func(name string) string {
		return fmt.Sprint("Hello there, ", name);
	}

	greeting2 := fgreet("Mary")

	fmt.Println(greeting2)
}