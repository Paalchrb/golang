package main

import "fmt"

func main() {
	printEveryThirdInRange(10, 35)
}

func printEveryThirdInRange(n int, m int) {
	for i := n ; i <= m ; i += 3 {
		fmt.Println(i)
	}
}