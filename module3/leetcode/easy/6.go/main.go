package main

import "fmt"

func uniqueMorseRepresentations(words []string) int {
	// Morse code representation of each letter
	morseCode := [26]string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.",
		"---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}

	// Set to store unique transformations
	transformations := make(map[string]bool)

	for _, word := range words {
		var morse string
		for _, char := range word {
			morse += morseCode[char-'a']
		}
		transformations[morse] = true
	}

	return len(transformations)
}

func main() {
	words1 := []string{"gin", "zen", "gig", "msg"}
	fmt.Println(uniqueMorseRepresentations(words1)) // Output: 2

	words2 := []string{"a"}
	fmt.Println(uniqueMorseRepresentations(words2)) // Output: 1
}
