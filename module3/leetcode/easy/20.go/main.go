package main

import "fmt"

func kidsWithCandies(candies []int, extraCandies int) []bool {
	// Находим максимальное количество конфет среди всех детей
	maxCandies := 0
	for _, candy := range candies {
		if candy > maxCandies {
			maxCandies = candy
		}
	}

	// Создаем массив для хранения результатов
	result := make([]bool, len(candies))

	// Проверяем, может ли каждый ребенок с extraCandies иметь больше или равное количество конфет, чем максимальное
	for i, candy := range candies {
		if candy+extraCandies >= maxCandies {
			result[i] = true
		}
	}

	return result
}

func main() {
	// Примеры использования
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3)) // Output: [true true true false true]
	fmt.Println(kidsWithCandies([]int{4, 2, 1, 1, 2}, 1)) // Output: [true false false false false]
	fmt.Println(kidsWithCandies([]int{12, 1, 12}, 10))    // Output: [true false true]
}
