package main

import "fmt"

func runningSum(nums []int) []int {
	runningSum := make([]int, len(nums))
	sum := 0
	for i, num := range nums {
		sum += num
		runningSum[i] = sum
	}
	return runningSum
}

func main() {
	nums1 := []int{1, 2, 3, 4}
	fmt.Println(runningSum(nums1)) // Output: [1 3 6 10]

	nums2 := []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(nums2)) // Output: [1 2 3 4 5]

	nums3 := []int{3, 1, 2, 10, 1}
	fmt.Println(runningSum(nums3)) // Output: [3 4 6 16 17]
}
