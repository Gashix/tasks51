package main

import "fmt"

func finalValueAfterOperations(operations []string) int {
	var X int
	for _, op := range operations {
		switch op {
		case "++X", "X++":
			X++
		case "--X", "X--":
			X--
		}
	}
	return X
}

func main() {
	operations1 := []string{"--X", "X++", "X++"}
	fmt.Println(finalValueAfterOperations(operations1)) // Output: 1

	operations2 := []string{"++X", "++X", "X++"}
	fmt.Println(finalValueAfterOperations(operations2)) // Output: 3

	operations3 := []string{"X++", "++X", "--X", "X--"}
	fmt.Println(finalValueAfterOperations(operations3)) // Output: 0
}
