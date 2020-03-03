package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(flipCase("Hello World"))
	fmt.Println(flipCase("HaHaHa"))
}

func flipCase(s string) string {
	var res string;
	for i := 0 ; i < len(s) ; i++ {
		if strings.ToLower(string(s[i])) == string(s[i]) {  //letter is lowercase
			res += strings.ToUpper(string(s[i]))
		} else {
			res += strings.ToLower(string(s[i]))
		}
	}
	return res
}