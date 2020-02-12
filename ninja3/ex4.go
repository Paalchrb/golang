package main

import "fmt"

func main() {
	i := 1987
	for {
		i++
		if(i > 2020) {
			break
		}
		fmt.Println(i)
	}
}