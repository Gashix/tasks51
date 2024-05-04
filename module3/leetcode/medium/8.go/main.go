package main

import "fmt"

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	incoming := make(map[int]bool)
	for _, edge := range edges {
		incoming[edge[1]] = true
	}

	result := make([]int, 0)
	for i := 0; i < n; i++ {
		if !incoming[i] {
			result = append(result, i)
		}
	}

	return result
}

func main() {
	n1 := 6
	edges1 := [][]int{{0, 1}, {0, 2}, {2, 5}, {3, 4}, {4, 2}}
	fmt.Println(findSmallestSetOfVertices(n1, edges1)) // Output: [0 3]

	n2 := 5
	edges2 := [][]int{{0, 1}, {2, 1}, {3, 1}, {1, 4}, {2, 4}}
	fmt.Println(findSmallestSetOfVertices(n2, edges2)) // Output: [0 2 3]
}
