package main

import "fmt"

func main() {
	for i := 1 ; i <= 6 ; i++ {
		var xi []int
		for j := 1 ; j <= 6 ; j++ {
			num := i * j
			xi = append(xi, num)
		}
		fmt.Println(xi)
	}
}
