package main

import "fmt"

func main() {
	printPattern(5, 3)
}

func printPattern(n int, m int) {
	for i := 0 ; i < n ; i++ {
		s := ""
		for j := 0 ; j < m ; j++ {
			s += "*"
		}
		fmt.Println(s)
	}
}