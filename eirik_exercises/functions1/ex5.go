package main

import "fmt"

func main() {
	printEveryDivisibleRange(10, 35, 3)
}

func printEveryDivisibleRange(n int, m int, x int) {
	for i := n ; i <= m ; i++ {
		if i % x == 0 {
			fmt.Println(i)
		}
	}
}