package main

import (
	"fmt"
	"sort"
)

//Person type
type Person struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByFirst implements sort.Interface for []Person based on the First field.
type ByFirst []Person

func (bf ByFirst) Len() int           { return len(bf) }
func (bf ByFirst) Swap(i, j int)      { bf[i], bf[j] = bf[j], bf[i] }
func (bf ByFirst) Less(i, j int) bool { return bf[i].First < bf[j].First }

func main() {
	u1 := Person{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := Person{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := Person{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	people := []Person{u1, u2, u3}

	fmt.Println(people)
	sort.Sort(ByAge(people))
	for _, person := range people{
		fmt.Println(person.Age, person.First)
	}
	
	sort.Sort(ByFirst(people))
	for _, person := range people{
		fmt.Println(person.First, person.Age)
	}
	
	for _, person := range people{
		sort.Strings(person.Sayings)
	}
	
	for i, person := range people {
		fmt.Println("Person #", i)
		fmt.Println("\t", person.First, person.Last, person.Age)
		for _, saying := range person.Sayings {
			fmt.Println("\t\t", saying)
		}
	}
}
