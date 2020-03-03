package main

import (
	"fmt"

	"github.com/GoesToEleven/go-programming/code_samples/010-ninja-level-thirteen/01/starting-code/dog"
)

type canine struct{
	name string
	age int
}

func main() {
	fido := canine{
		name: "fido",
		age: dog.Years(7),
	}

	fmt.Println(fido.age)
}