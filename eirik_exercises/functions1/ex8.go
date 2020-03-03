package main

import "fmt"

func main() {
	printPattern(7)
}

func printPattern(n int) {
	s := ""
	for i := 0 ; i < n ; i++ {
		s += "*"
		fmt.Println(s)
	}
	for j := n - 1 ; j > 0 ; j-- {
		fmt.Println(s[0: j])
	}
}

