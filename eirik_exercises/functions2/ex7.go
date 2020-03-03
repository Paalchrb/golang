package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(areaOfCircle(4))
	fmt.Println(areaOfCircle(3.2131))
}

func areaOfCircle(r float64) float64 {
	return math.Pi * r * r
}