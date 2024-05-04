package main

import (
	"fmt"
	"strings"
)

func mostWordsFound(sentences []string) int {
	max := 0
	for _, sentence := range sentences {
		words := strings.Split(sentence, " ")
		count := len(words)
		if count > max {
			max = count
		}
	}
	return max
}

func main() {
	sentences1 := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	fmt.Println(mostWordsFound(sentences1)) // Output: 6

	sentences2 := []string{"please wait", "continue to fight", "continue to win"}
	fmt.Println(mostWordsFound(sentences2)) // Output: 3
}
