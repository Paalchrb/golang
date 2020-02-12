package main

import "fmt"

type vehicle struct{
	doors int
	color string
}

type sedan struct{
	vehicle
	luxury bool
}

type truck struct{
	vehicle
	fourWheel bool
}

func main() {
	carSedan := sedan{
		vehicle: vehicle{
			doors: 5,
			color: "red",
		},
		luxury: true,
	}

	carTruck := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheel: true,
	}

	fmt.Println(carSedan)
	fmt.Println("-----------")
	fmt.Println(carTruck)

	fmt.Println(carSedan.color)
	fmt.Println(carTruck.doors)
}