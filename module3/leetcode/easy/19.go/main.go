package main

import (
	"fmt"
	"sort"
	"strconv"
)

func minimumSum(num int) int {
	// Преобразуем число в строку для удобства работы с цифрами
	numStr := strconv.Itoa(num)

	// Создаем срез для хранения цифр
	digits := make([]int, len(numStr))

	// Заполняем срез цифрами из числа
	for i, char := range numStr {
		digit, _ := strconv.Atoi(string(char))
		digits[i] = digit
	}

	// Сортируем цифры по возрастанию
	sort.Ints(digits)

	// Создаем два новых числа
	var new1, new2 int

	// Распределяем цифры между new1 и new2
	for i, digit := range digits {
		if i%2 == 0 {
			new1 = new1*10 + digit
		} else {
			new2 = new2*10 + digit
		}
	}

	// Возвращаем минимально возможную сумму new1 и new2
	return new1 + new2
}

func main() {
	// Примеры использования
	fmt.Println(minimumSum(2932)) // Output: 52
	fmt.Println(minimumSum(4009)) // Output: 13
}
