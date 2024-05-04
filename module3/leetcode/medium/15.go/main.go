package main

import (
	"fmt"
	"strconv"
)

func minPartitions(n string) int {
	maxDigit := 0
	for _, char := range n {
		digit, _ := strconv.Atoi(string(char))
		if digit > maxDigit {
			maxDigit = digit
		}
	}
	return maxDigit
}

func main() {
	fmt.Println(minPartitions("32"))                   // Output: 3
	fmt.Println(minPartitions("82734"))                // Output: 8
	fmt.Println(minPartitions("27346209830709182346")) // Output: 9
}
