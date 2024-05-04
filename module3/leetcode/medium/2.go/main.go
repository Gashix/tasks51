package main

import (
	"fmt"
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] > score[j][k]
	})
	return score
}

func main() {
	score1 := [][]int{{10, 6, 9, 1}, {7, 5, 11, 2}, {4, 8, 3, 15}}
	k1 := 2
	fmt.Println(sortTheStudents(score1, k1))
	// Output: [[7 5 11 2] [10 6 9 1] [4 8 3 15]]

	score2 := [][]int{{3, 4}, {5, 6}}
	k2 := 0
	fmt.Println(sortTheStudents(score2, k2))
	// Output: [[5 6] [3 4]]
}
