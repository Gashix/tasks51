package main

import "fmt"

func numJewelsInStones(jewels string, stones string) int {
	stoneCount := make(map[rune]int)
	for _, stone := range stones {
		stoneCount[stone]++
	}

	jewelCount := 0
	for _, jewel := range jewels {
		jewelCount += stoneCount[jewel]
	}

	return jewelCount
}

func main() {
	jewels1 := "aA"
	stones1 := "aAAbbbb"
	fmt.Println(numJewelsInStones(jewels1, stones1)) // Output: 3

	jewels2 := "z"
	stones2 := "ZZ"
	fmt.Println(numJewelsInStones(jewels2, stones2)) // Output: 0
}
