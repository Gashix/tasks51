package main

import (
	"fmt"
	"math"
)

func countGoodTriplets(arr []int, a, b, c int) int {
	count := 0
	n := len(arr)

	// Iterate through all possible triplets (i, j, k)
	for i := 0; i < n-2; i++ {
		for j := i + 1; j < n-1; j++ {
			for k := j + 1; k < n; k++ {
				// Check if the triplet satisfies all conditions
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) &&
					math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
					math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(countGoodTriplets([]int{3, 0, 1, 1, 9, 7}, 7, 2, 3)) // Output: 4
	fmt.Println(countGoodTriplets([]int{1, 1, 2, 2, 3}, 0, 0, 1))    // Output: 0
}
