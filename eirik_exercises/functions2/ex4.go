package main

import "fmt"

func main() {
	fmt.Println(repeat("hello", 5))
	fmt.Println(repeat("hel te seee", 10))
}

func repeat(s string, n int) string {
	var str string
	for i := 0 ; i < n ; i++ {
		str += s
	}
	return str
}