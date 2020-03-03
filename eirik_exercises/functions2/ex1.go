package main

import (
	"fmt"
	"strconv"
)

func main() {
	res := isPalindrome(1234321)
	fmt.Println(res)
	res = isPalindrome(456654)
	fmt.Println(res)
	res = isPalindrome(456123654)
	fmt.Println(res)
}

func isPalindrome(n int) bool {
	s:= strconv.Itoa(n)									//convert num to string
	fmt.Printf("%q\n", s)

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



