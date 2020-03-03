package main

import "fmt"

func main() {
	list1 := []int{5, 4, 3, 2, 1}
	list2 := []int{3, 5, 432 ,122, 12, 314, 1, 9, 21, 13, 8}
	
	fmt.Println(sortList(list1))
	fmt.Println(sortList(list2))
}

func sortList(xi []int) []int {
	for {
		swapped := false
		for i := 0 ; i < len(xi) - 1  ; i++ {
			if xi[i] > xi[i + 1] {
				xi[i], xi[i + 1] = xi[i + 1], xi[i]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return xi
}
