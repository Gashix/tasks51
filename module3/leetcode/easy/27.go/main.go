package main

import "fmt"

func balancedStringSplit(s string) int {
	balance := 0
	count := 0

	for _, char := range s {
		if char == 'L' {
			balance++
		} else {
			balance--
		}

		// If balance becomes zero, increment the count of balanced strings
		if balance == 0 {
			count++
		}
	}

	return count
}

func main() {
	// Test cases
	fmt.Println(balancedStringSplit("RLRRLLRLRL")) // Output: 4
	fmt.Println(balancedStringSplit("RLRRRLLRLL")) // Output: 2
	fmt.Println(balancedStringSplit("LLLLRRRR"))   // Output: 1
}
