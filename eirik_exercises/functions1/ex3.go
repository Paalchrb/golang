package main

import "fmt"

func main() {
	printEveryXInRange(10, 35, 5)
}

func printEveryXInRange(n int, m int, x int) {
	for i := n ; i <= m ; i += x {
		fmt.Println(i)
	}
}