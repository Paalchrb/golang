package main

import "fmt"

func main() {
	i := 64
	for {
		i++
		if(i >= 91) {
			break
		}
		fmt.Println(i)
		for j := 0 ; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}