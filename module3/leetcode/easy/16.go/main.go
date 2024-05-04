package main

import "fmt"

func smallestEvenMultiple(n int) int {
	if n%2 == 0 {
		return n
	}
	return 2 * n
}

func main() {
	fmt.Println(smallestEvenMultiple(5)) // Output: 10
	fmt.Println(smallestEvenMultiple(6)) // Output: 6
}
