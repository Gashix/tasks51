package main

import "fmt"

func createTargetArray(nums []int, index []int) []int {
	target := make([]int, 0, len(nums))

	for i := 0; i < len(nums); i++ {
		target = insertAtIndex(target, index[i], nums[i])
	}

	return target
}

func insertAtIndex(arr []int, index, value int) []int {
	arr = append(arr, value)
	copy(arr[index+1:], arr[index:])
	arr[index] = value

	return arr
}

func main() {
	// Test cases
	fmt.Println(createTargetArray([]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1})) // Output: [0 4 1 3 2]
	fmt.Println(createTargetArray([]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0})) // Output: [0 1 2 3 4]
	fmt.Println(createTargetArray([]int{1}, []int{0}))                         // Output: [1]
}
