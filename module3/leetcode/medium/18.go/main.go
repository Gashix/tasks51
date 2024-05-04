package main

import "fmt"

func groupThePeople(groupSizes []int) [][]int {
	groups := make(map[int][]int)
	result := make([][]int, 0)

	for i, size := range groupSizes {
		groups[size] = append(groups[size], i)
	}

	for size, group := range groups {
		for i := 0; i < len(group); i += size {
			result = append(result, group[i:i+size])
		}
	}

	return result
}

func main() {
	groupSizes := []int{3, 3, 3, 3, 3, 1, 3}
	fmt.Println("Input:", groupSizes)
	fmt.Println("Output:", groupThePeople(groupSizes))

	groupSizes = []int{2, 1, 3, 3, 3, 2}
	fmt.Println("Input:", groupSizes)
	fmt.Println("Output:", groupThePeople(groupSizes))
}
