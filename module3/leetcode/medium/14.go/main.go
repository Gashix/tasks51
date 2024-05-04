package main

import "fmt"

func xorQueries(arr []int, queries [][]int) []int {
	n := len(arr)
	xor := make([]int, n+1)

	// Precompute XOR of all elements in arr
	for i := 1; i <= n; i++ {
		xor[i] = xor[i-1] ^ arr[i-1]
	}

	// Compute XOR of elements in queries
	result := make([]int, len(queries))
	for i, q := range queries {
		result[i] = xor[q[1]+1] ^ xor[q[0]]
	}

	return result
}

func main() {
	arr := []int{1, 3, 4, 8}
	queries := [][]int{{0, 1}, {1, 2}, {0, 3}, {3, 3}}
	fmt.Println(xorQueries(arr, queries)) // Output: [2 7 14 8]

	arr = []int{4, 8, 2, 10}
	queries = [][]int{{2, 3}, {1, 3}, {0, 0}, {0, 3}}
	fmt.Println(xorQueries(arr, queries)) // Output: [8 0 4 4]
}
