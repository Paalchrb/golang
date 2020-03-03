package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(capitalize("helloworld"))
	fmt.Println(capitalize("hELloWoRLd"))
}

func capitalize(s string) string {
	s = strings.ToLower(s)		//lowercase all letters
	s = strings.Title(s) 			//uppercase first letter
	
	return s
}