package main

import (
	"fmt"
)

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (sq *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			sq.rectangle[i][j] = newValue
		}
	}
}

func (sq *SubrectangleQueries) GetValue(row int, col int) int {
	return sq.rectangle[row][col]
}

func main() {
	rect := [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	obj := Constructor(rect)
	fmt.Println(obj.GetValue(0, 0)) // Output: 1
	obj.UpdateSubrectangle(0, 0, 2, 2, 100)
	fmt.Println(obj.GetValue(0, 0)) // Output: 100
	fmt.Println(obj.GetValue(2, 2)) // Output: 100
}
