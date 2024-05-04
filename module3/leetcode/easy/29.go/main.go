package main

import "fmt"

func xorOperation(n, start int) int {
	result := start

	for i := 1; i < n; i++ {
		// Calculate each element nums[i] = start + 2 * i and perform XOR with result
		result ^= start + 2*i
	}

	return result
}

func main() {
	// Test cases
	fmt.Println(xorOperation(5, 0)) // Output: 8
	fmt.Println(xorOperation(4, 3)) // Output: 8
}
