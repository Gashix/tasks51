package main

import "fmt"

func buildArray(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func main() {
	nums1 := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(nums1)) // Output: [0 1 2 4 5 3]

	nums2 := []int{5, 0, 1, 2, 3, 4}
	fmt.Println(buildArray(nums2)) // Output: [4 5 0 1 2 3]
}
