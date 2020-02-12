package main

import "fmt"

func main() {
	xi := []int{1,2,3,4,5,6,7,8,9}
	sumEven := sumEvenNumbers(sum, xi...)
	fmt.Println(sumEven)
}

func sumEvenNumbers(f func(x ...int) int, y ...int) int {
	var xi []int
	for _, v := range y {
		if v % 2 == 0 {
			xi = append(xi, v)
		}
	}
	return f(xi...)
}

func sum(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}