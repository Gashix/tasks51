package main

import (
	"fmt"
)

func maxSum(grid [][]int) int {
	maxSum := -1 << 31 // Initialize maxSum to the smallest possible integer
	m, n := len(grid), len(grid[0])

	// Iterate through each cell in the grid
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// Compute the sum of the hourglass centered at (i, j)
			sum := grid[i-1][j-1] + grid[i-1][j] + grid[i-1][j+1] +
				grid[i][j] +
				grid[i+1][j-1] + grid[i+1][j] + grid[i+1][j+1]

			// Update maxSum if necessary
			if sum > maxSum {
				maxSum = sum
			}
		}
	}

	return maxSum
}

func main() {
	grid1 := [][]int{{6, 2, 1, 3}, {4, 2, 1, 5}, {9, 2, 8, 7}, {4, 1, 2, 9}}
	fmt.Println("Max Hourglass Sum (Example 1):", maxSum(grid1))

	grid2 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("Max Hourglass Sum (Example 2):", maxSum(grid2))
}
