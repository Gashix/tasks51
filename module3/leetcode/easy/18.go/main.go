package main

import (
	"fmt"
)

func differenceOfSum(nums []int) int {
	elementSum := 0
	digitSum := 0

	// Calculate the element sum
	for _, num := range nums {
		elementSum += num
	}

	// Calculate the digit sum
	for _, num := range nums {
		for num > 0 {
			digitSum += num % 10
			num /= 10
		}
	}

	// Calculate the absolute difference
	diff := elementSum - digitSum
	if diff < 0 {
		diff = -diff
	}

	return diff
}

func main() {
	nums1 := []int{1, 15, 6, 3}
	fmt.Println(differenceOfSum(nums1)) // Output: 9

	nums2 := []int{1, 2, 3, 4}
	fmt.Println(differenceOfSum(nums2)) // Output: 0
}
