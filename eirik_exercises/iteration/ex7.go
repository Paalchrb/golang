package main

import "fmt"

func main() {
	for i := -20 ; i <= 20 ; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}
