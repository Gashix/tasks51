package main

import "fmt"

func convertTemperature(celsius float64) []float64 {
	kelvin := celsius + 273.15
	fahrenheit := celsius*1.80 + 32.00
	return []float64{kelvin, fahrenheit}
}

func main() {
	celsius1 := 36.50
	fmt.Println(convertTemperature(celsius1)) // Output: [309.65 97.7]

	celsius2 := 122.11
	fmt.Println(convertTemperature(celsius2)) // Output: [395.26 251.798]
}
