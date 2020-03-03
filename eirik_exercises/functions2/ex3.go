package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(wordCount("Hello world how many words am I wubba dubba dub dub dubdubdub?"))
	fmt.Println(wordCount("This is test sentence number two, wohoooo!   "))
}

func wordCount(s string) int {
	s = strings.TrimSpace(s)		//removes whitespace before and after string
	words := 0

	for i := 0 ; i < len(s) ; i++ {
		if string(s[i]) == " " {
			words++
		}
	}
	return words + 1
}