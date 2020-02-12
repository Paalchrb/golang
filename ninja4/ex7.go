package main

import "fmt"

func main() {
	james := []string{"James", "Bond", "Shaken, not stirred"}
	money := []string{"Miss", "Moneypenny", "Hellooooo, James"}
	fmt.Println(james)
	fmt.Println(money)

	people := [][]string{james, money}
	fmt.Println(people)

	for i := 0 ; i < len(people) ; i++ {
		fmt.Printf("People index: %v\n", i)
		for j := 0 ; j < len(people[i]) ; j++ {
			fmt.Printf("\t%v\n", people[i][j])
		}
	}
}