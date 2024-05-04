package main

import "fmt"

func defangIPaddr(address string) string {
	var defanged string
	for _, char := range address {
		if char == '.' {
			defanged += "[.]"
		} else {
			defanged += string(char)
		}
	}
	return defanged
}

func main() {
	address1 := "1.1.1.1"
	fmt.Println(defangIPaddr(address1)) // Output: "1[.]1[.]1[.]1"

	address2 := "255.100.50.0"
	fmt.Println(defangIPaddr(address2)) // Output: "255[.]100[.]50[.]0"
}
