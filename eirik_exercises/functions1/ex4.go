package main

import "fmt"

func main() {
	printEveryXInRangeBackwards(10, 35, 3)
}

func printEveryXInRangeBackwards(n int, m int, x int) {
	for i := m ; i >= n ; i -= x {
		fmt.Println(i)
	}
}