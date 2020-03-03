package main

import (
	"fmt"
)

func main() {
	list1 := []int{1, 1, 2, 2, 3, 4, 4, 5, 5, 9}
	list2 := []int{1, 2, 2, 2, 3}

	fmt.Println(removeDuplicates(list1))
	fmt.Println(removeDuplicates(list2))
}

func removeDuplicates(xi []int) []int {
	var dupesList []int
	for _, v := range xi {
		if firstIndex(xi, v) != lastIndex(xi, v) && firstIndex(dupesList, v) == -1 {
			dupesList = append(dupesList, v)
		}
	}
	return dupesList
}

func firstIndex(slice []int, val int) int {
	for i:= 0 ; i < len(slice); i++ {
			if slice[i] == val {
					return i
			}
	}
	return -1
}

func lastIndex (slice []int, val int) int {
	for i:= len(slice) - 1 ; i >= 0 ; i-- {
		if slice[i] == val {
				return i
		}
}
return -1
}