package main

import (
	"fmt"
	"math"
)

type square struct{
	length float64
}

type circle struct{
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface{
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circle := circle{
		radius: 4.123,
	}

	square := square{
		length: 6.132,
	}

	info(circle)
	info(square)
}