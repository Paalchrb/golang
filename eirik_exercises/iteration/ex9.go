package main

import "fmt"

func main() {
	s := ""
	for i := 0 ; i < 4 ; i++ {
		s += "*"
		fmt.Println(s)
	}
	for j := 3 ; j > 0 ; j-- {
		fmt.Println(s[0: j])
	}
}
