package main

import "fmt"

func main() {
	for i := 0 ; i <= 20 ; i++ {
		if i == 13 || i == 17 {
			continue
		}
		fmt.Println(i)
	}
}
