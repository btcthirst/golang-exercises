package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	people := []Person{
		{"Alice", 30},
		{"Bob", 25},
		{"Charlie", 35},
		{"David", 20},
		{"Eve", 28},
		{"Frank", 22},
		{"Grace", 27},
		{"Heidi", 22},
	}

	// Sort by age
	sortByAge(people)
	fmt.Println("Sorted by age:")

	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}

	// Sort by name
	sortByName(people)
	fmt.Println("\nSorted by name:")
	for _, person := range people {
		fmt.Println(person.Name, person.Age)
	}
}

func sortByAge(people []Person) {
	// Sort by age
	sort.Slice(people, func(i, j int) bool {
		if people[i].Age != people[j].Age {
			return people[i].Age < people[j].Age
		}
		return people[i].Name < people[j].Name
	})
}

func sortByName(people []Person) {
	// Sort by name
	sort.Slice(people, func(i, j int) bool {
		if people[i].Name != people[j].Name {
			return people[i].Name < people[j].Name
		}
		return people[i].Age < people[j].Age
	})
}
