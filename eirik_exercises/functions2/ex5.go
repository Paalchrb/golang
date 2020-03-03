package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("I am a happy dog"))
}

func reverseWords(s string) string {
	words := strings.Split(s, " ")
	result := ""

	for _, v := range words {
		for i := len(v) - 1 ; i >= 0 ; i-- {
			result += string(v[i])
		}
		result += " "
	}

	return result
}