package main

import "fmt"

func numTilePossibilities(tiles string) int {
	freq := make([]int, 26)
	for _, ch := range tiles {
		freq[ch-'A']++
	}
	return dfs(freq)
}

func dfs(freq []int) int {
	sum := 0
	for i := 0; i < 26; i++ {
		if freq[i] == 0 {
			continue
		}
		sum++
		freq[i]--
		sum += dfs(freq)
		freq[i]++
	}
	return sum
}

func main() {
	tiles1 := "AAB"
	fmt.Println(numTilePossibilities(tiles1)) // Output: 8

	tiles2 := "AAABBC"
	fmt.Println(numTilePossibilities(tiles2)) // Output: 188

	tiles3 := "V"
	fmt.Println(numTilePossibilities(tiles3)) // Output: 1
}
