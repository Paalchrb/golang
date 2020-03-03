package main

import (
	"fmt"
	"strings"
)

func main() {
	res := isPalindrome(" Rats li ve oN NO eviLStaR")
	fmt.Println(res)
	res = isPalindrome(" raceCar")
	fmt.Println(res)
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)								//lowercase all letters
	s = strings.ReplaceAll(s, " ", "")		//removes whitespace
	left := 0
	right := len(s) - 1

	for left < right {									//while left pointer is less than right pointer
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}



