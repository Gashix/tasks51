package main

import "fmt"

func interpret(command string) string {
	result := ""
	i := 0
	for i < len(command) {
		if command[i] == 'G' {
			result += "G"
			i++
		} else if command[i] == '(' && command[i+1] == ')' {
			result += "o"
			i += 2
		} else if command[i] == '(' && command[i+1] == 'a' {
			result += "al"
			i += 4
		}
	}
	return result
}

func main() {
	// Test cases
	fmt.Println(interpret("G()(al)"))        // Output: "Goal"
	fmt.Println(interpret("G()()()()(al)"))  // Output: "Gooooal"
	fmt.Println(interpret("(al)G(al)()()G")) // Output: "alGalooG"
}
