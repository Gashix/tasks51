package main

import "fmt"

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	result := make([]bool, len(l))

	for i, left := range l {
		right := r[i]
		subarray := nums[left : right+1]

		elements := make(map[int]bool)
		for _, num := range subarray {
			elements[num] = true
		}

		minNum := subarray[0]
		maxNum := subarray[0]
		for _, num := range subarray {
			if num < minNum {
				minNum = num
			}
			if num > maxNum {
				maxNum = num
			}
		}

		diff := (maxNum - minNum) / (right - left)

		isArithmetic := true
		for j := minNum; j <= maxNum; j += diff {
			if !elements[j] {
				isArithmetic = false
				break
			}
		}
		result[i] = isArithmetic
	}

	return result
}
func main() {
	nums1 := []int{4, 6, 5, 9, 3, 7}
	l1 := []int{0, 0, 2}
	r1 := []int{2, 3, 5}
	fmt.Println(checkArithmeticSubarrays(nums1, l1, r1)) // Output: [true false true]

	nums2 := []int{-12, -9, -3, -12, -6, 15, 20, -25, -20, -15, -10}
	l2 := []int{0, 1, 6, 4, 8, 7}
	r2 := []int{4, 4, 9, 7, 9, 10}
	fmt.Println(checkArithmeticSubarrays(nums2, l2, r2)) // Output: [false true false false true true]
}
