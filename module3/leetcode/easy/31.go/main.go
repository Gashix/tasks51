package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name   string
	Height int
}

func sortPeople(names []string, heights []int) []string {
	n := len(names)
	people := make([]Person, n)

	for i := 0; i < n; i++ {
		people[i] = Person{Name: names[i], Height: heights[i]}
	}

	sort.Slice(people, func(i, j int) bool {
		return people[i].Height > people[j].Height
	})

	sortedNames := make([]string, n)
	for i := 0; i < n; i++ {
		sortedNames[i] = people[i].Name
	}

	return sortedNames
}

func main() {
	fmt.Println(sortPeople([]string{"Mary", "John", "Emma"}, []int{180, 165, 170})) // Output: ["Mary","Emma","John"]
	fmt.Println(sortPeople([]string{"Alice", "Bob", "Bob"}, []int{155, 185, 150}))  // Output: ["Bob","Alice","Bob"]
}
