package main

import "fmt"

func main() {
	xi := []int{1,2,3,4,5,6,7,8,9}
	sum := foo(xi...)
	fmt.Println(sum)
	sum2 := bar(xi)
	fmt.Println(sum2)

}

func foo(x ...int) int {
	var total int
	for _, v := range x {
		total += v
	}
	return total
}

func bar(xii []int) int {
	var total int
	for _, v := range xii {
		total += v
	}
	return total
}