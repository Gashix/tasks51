package main

import "fmt"

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)
	for i := 0; i < n; i++ {
		ans[i] = nums[i]
		ans[i+n] = nums[i]
	}
	return ans
}

func main() {
	nums1 := []int{1, 2, 1}
	fmt.Println(getConcatenation(nums1)) // Output: [1 2 1 1 2 1]

	nums2 := []int{1, 3, 2, 1}
	fmt.Println(getConcatenation(nums2)) // Output: [1 3 2 1 1 3 2 1]
}
