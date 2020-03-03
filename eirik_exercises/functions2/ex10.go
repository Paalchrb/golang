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
	var filteredList []int
	for _, value := range xi {
		exists := findItem(filteredList, value)
		if exists {
			continue
		}
		filteredList = append(filteredList, value)
	}
	return filteredList
}

func findItem(slice []int, val int) (bool) {
	for _, item := range slice {
			if item == val {
					return true
			}
	}
	return false
}