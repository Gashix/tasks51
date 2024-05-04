package main

import "fmt"

func numIdenticalPairs(nums []int) int {
	count := make(map[int]int)
	for _, num := range nums {
		count[num]++
	}

	goodPairs := 0
	for _, c := range count {
		goodPairs += c * (c - 1) / 2
	}

	return goodPairs
}

func main() {
	nums1 := []int{1, 2, 3, 1, 1, 3}
	fmt.Println(numIdenticalPairs(nums1)) // Output: 4

	nums2 := []int{1, 1, 1, 1}
	fmt.Println(numIdenticalPairs(nums2)) // Output: 6

	nums3 := []int{1, 2, 3}
	fmt.Println(numIdenticalPairs(nums3)) // Output: 0
}
