package main

import "fmt"

func smallerNumbersThanCurrent(nums []int) []int {
	// Initialize a hashmap to count occurrences of each number
	count := make(map[int]int)

	// Count occurrences of each number
	for _, num := range nums {
		count[num]++
	}

	// Iterate through the array to count smaller numbers
	result := make([]int, len(nums))
	for i, num := range nums {
		// Initialize count of smaller numbers as 0
		smaller := 0
		// Iterate through numbers smaller than num
		for j := 0; j < num; j++ {
			smaller += count[j]
		}
		// Update the result
		result[i] = smaller
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(smallerNumbersThanCurrent([]int{8, 1, 2, 2, 3})) // Output: [4 0 1 1 3]
	fmt.Println(smallerNumbersThanCurrent([]int{6, 5, 4, 8}))    // Output: [2 1 0 3]
	fmt.Println(smallerNumbersThanCurrent([]int{7, 7, 7, 7}))    // Output: [0 0 0 0]
}
