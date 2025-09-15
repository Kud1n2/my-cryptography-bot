package main

import (
	"fmt"
)

func binaryShtuka(key, number int) (string, int) {
	bynKey := fmt.Sprintf("%b", key)
	return bynKey, len(bynKey)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func encryption(key, number, module int) int {
	if module == 0 {
		return 0 // Защита от деления на ноль
	}

	binarnik, length := binaryShtuka(key, number)
	numbers := make([]int, length)

	// Инициализация первого элемента
	numbers[0] = number

	// Заполнение массива numbers
	for i := 1; i < length; i++ {
		numbers[i] = (numbers[i-1] * numbers[i-1]) % module
	}

	multiply := 1
	// Проход по бинарному представлению

	binarnik = Reverse(binarnik)
	for i, edinichka := range binarnik {
		if edinichka == '1' {
			// Безопасный доступ к индексу
			if i >= 0 && i < length {
				multiply = (multiply * numbers[i]) % module
			}
		}
	}

	return multiply
}
