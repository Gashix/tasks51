package main

import "fmt"

func decompressRLElist(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i += 2 {
		freq := nums[i]
		val := nums[i+1]

		// Repeat val freq times and append to the result
		for j := 0; j < freq; j++ {
			result = append(result, val)
		}
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(decompressRLElist([]int{1, 2, 3, 4})) // Output: [2 4 4 4]
	fmt.Println(decompressRLElist([]int{1, 1, 2, 3})) // Output: [1 3 3]
}
