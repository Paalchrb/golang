package main

import "fmt"

func main() {
	printTable(4, 4)
}

func printTable(n int, m int) {
	for i := 1 ; i <= n ; i++ {
		var xi []int
		for j := 1 ; j <= m ; j++ {
			num := i * j
			xi = append(xi, num)
		}
		fmt.Println(xi)
	}
}
