package main

import "fmt"

func main() {
	printPattern(2)
}

func printPattern(n int) {
	for i := 0 ; i < n ; i++ {
		s := ""
		for j := 0 ; j < n ; j++ {
			s += "*"
		}
		fmt.Println(s)
	}
}