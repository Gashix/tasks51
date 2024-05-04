package main

import (
	"fmt"
	"strconv"
)

func subtractProductAndSum(n int) int {
	// Convert the integer to a string to extract its digits
	digits := strconv.Itoa(n)

	// Initialize variables to store product and sum
	product := 1
	sum := 0

	// Iterate through each digit
	for _, digit := range digits {
		// Convert the digit character back to integer
		d, _ := strconv.Atoi(string(digit))
		// Update product and sum
		product *= d
		sum += d
	}

	// Calculate the difference between product and sum
	return product - sum
}

func main() {
	// Test cases
	fmt.Println(subtractProductAndSum(234))  // Output: 15
	fmt.Println(subtractProductAndSum(4421)) // Output: 21
}
