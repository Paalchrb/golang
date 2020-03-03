package main

import "fmt"

func main() {
	printRange(10, 35)
}

func printRange(n int, m int) {
	for i := n ; i <= m ; i++ {
		fmt.Println(i)
	}
}