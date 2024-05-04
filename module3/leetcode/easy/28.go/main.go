package main

import (
	"fmt"
	"strconv"
)

func countDigits(num int) int {
	count := 0

	// Convert num to string to iterate through its digits
	numStr := strconv.Itoa(num)

	for _, digitChar := range numStr {
		digit, _ := strconv.Atoi(string(digitChar))

		// Check if the digit divides num
		if digit != 0 && num%digit == 0 {
			count++
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(countDigits(7))    // Output: 1
	fmt.Println(countDigits(121))  // Output: 2
	fmt.Println(countDigits(1248)) // Output: 4
}
