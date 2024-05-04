package main

import "fmt"

func decode(encoded []int, first int) []int {
	n := len(encoded) + 1
	arr := make([]int, n)
	arr[0] = first

	// Decoding using XOR property
	for i := 1; i < n; i++ {
		arr[i] = arr[i-1] ^ encoded[i-1]
	}

	return arr
}

func main() {
	// Test cases
	fmt.Println(decode([]int{1, 2, 3}, 1))    // Output: [1 0 2 1]
	fmt.Println(decode([]int{6, 2, 7, 3}, 4)) // Output: [4 2 0 7 4]
}
