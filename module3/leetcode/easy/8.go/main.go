package main

import "fmt"

func findKthPositive(arr []int, k int) int {
	missing := make(map[int]bool)

	for _, num := range arr {
		missing[num] = true
	}

	count := 0
	for i := 1; ; i++ {
		if !missing[i] {
			count++
			if count == k {
				return i
			}
		}
	}
}

func main() {
	arr1 := []int{2, 3, 4, 7, 11}
	k1 := 5
	fmt.Println(findKthPositive(arr1, k1)) // Output: 9

	arr2 := []int{1, 2, 3, 4}
	k2 := 2
	fmt.Println(findKthPositive(arr2, k2)) // Output: 6
}
