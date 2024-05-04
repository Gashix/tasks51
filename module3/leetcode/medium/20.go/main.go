package main

import "fmt"

func processQueries(queries []int, m int) []int {
	// Initialize the permutation P and the positions map
	P := make([]int, m)
	positions := make(map[int]int)
	for i := 0; i < m; i++ {
		P[i] = i + 1
		positions[i+1] = i
	}

	// Process each query
	result := make([]int, len(queries))
	for i, q := range queries {
		// Find the position of the query in the permutation
		pos := positions[q]
		result[i] = pos

		// Move the query to the beginning of the permutation
		for j := pos; j > 0; j-- {
			P[j], P[j-1] = P[j-1], P[j]
			positions[P[j]] = j
			positions[P[j-1]] = j - 1
		}
	}

	return result
}

func main() {
	// Example usage:
	queries1 := []int{3, 1, 2, 1}
	m1 := 5
	fmt.Println(processQueries(queries1, m1)) // Output: [2 1 2 1]

	queries2 := []int{4, 1, 2, 2}
	m2 := 4
	fmt.Println(processQueries(queries2, m2)) // Output: [3 1 2 0]

	queries3 := []int{7, 5, 5, 8, 3}
	m3 := 8
	fmt.Println(processQueries(queries3, m3)) // Output: [6 5 0 7 5]
}
