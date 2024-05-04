package main

import "fmt"

func shuffle(nums []int, n int) []int {
	shuffled := make([]int, len(nums))

	for i := 0; i < n; i++ {
		shuffled[2*i] = nums[i]
		shuffled[2*i+1] = nums[n+i]
	}

	return shuffled
}

func main() {
	nums1 := []int{2, 5, 1, 3, 4, 7}
	n1 := 3
	fmt.Println(shuffle(nums1, n1)) // Output: [2 3 5 4 1 7]

	nums2 := []int{1, 2, 3, 4, 4, 3, 2, 1}
	n2 := 4
	fmt.Println(shuffle(nums2, n2)) // Output: [1 4 2 3 3 2 4 1]

	nums3 := []int{1, 1, 2, 2}
	n3 := 2
	fmt.Println(shuffle(nums3, n3)) // Output: [1 2 1 2]
}
