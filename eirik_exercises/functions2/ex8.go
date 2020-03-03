package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(exponent(4, 5))
	fmt.Println(exponent(2, 8))
	fmt.Println(exponent(8, 3))
}

func exponent(n float64, exp float64) float64 {
	return math.Pow(n, exp)
}