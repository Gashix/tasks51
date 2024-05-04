package main

import "fmt"

func maximumWealth(accounts [][]int) int {
	maxWealth := 0

	for _, customer := range accounts {
		wealth := 0
		for _, money := range customer {
			wealth += money
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}

	return maxWealth
}

func main() {
	accounts1 := [][]int{{1, 2, 3}, {3, 2, 1}}
	fmt.Println(maximumWealth(accounts1)) // Output: 6

	accounts2 := [][]int{{1, 5}, {7, 3}, {3, 5}}
	fmt.Println(maximumWealth(accounts2)) // Output: 10

	accounts3 := [][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}
	fmt.Println(maximumWealth(accounts3)) // Output: 17
}
