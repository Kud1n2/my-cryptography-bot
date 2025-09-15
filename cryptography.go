package main

import (
	"fmt"
	"math"
)

func checkSimpleNumber(x int) bool {
	for i := 2; i < int(math.Sqrt(float64(x)))+1; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
func findSimpleNumbers(module int) (int, int) {
	if checkSimpleNumber(module) {
		fmt.Print("Невозможно разложить, т.к. число простое")
		return 0, 0
	}
	for i := 2; i < int(math.Sqrt(float64(module)))+1; i++ {
		if module%i == 0 {
			if checkSimpleNumber(i) && checkSimpleNumber(module/i) {
				//fmt.Print("Разложение на 2 простых числа: ")
				return i, module / i
			}
		}
	}
	fmt.Print("Невозможно разложить")
	return 0, 0
}

func fiN(module int) int {
	var x1, x2 int
	x1, x2 = findSimpleNumbers(module)
	fiN1 := module * (x1 - 1) * (x2 - 1)
	fiN2 := fiN1 / (x1 * x2)
	//fmt.Print("Фи(н)")
	return int(fiN2)
}

func gausMethod(num1, num2 int) int {
	var matrix [2][3]int = [2][3]int{{1, 0, num2}, {0, 1, num1}}
	for matrix[0][2] != 0 && matrix[1][2] != 0 {
		// Находим индексы строк с максимальным и минимальным значениями в третьем столбце
		maxIdx, minIdx := 0, 1
		if matrix[0][2] < matrix[1][2] {
			maxIdx, minIdx = 1, 0
		}

		// Вычисляем коэффициент для вычитания
		coefficient := matrix[maxIdx][2] / matrix[minIdx][2]

		// Вычитаем из строки с большим значением строку с меньшим, умноженную на коэффициент
		for i := 0; i < 3; i++ {
			matrix[maxIdx][i] -= coefficient * matrix[minIdx][i]
		}

		//fmt.Printf("\nПосле вычитания (коэффициент %d):\n", coefficient)
		//printMatrix(matrix)
	}
	if matrix[0][2] == 1 {
		if matrix[0][1] < 0 {
			for i := 0; i < 3; i++ {
				matrix[0][i] -= (-1) * matrix[1][i]
			}
		}
		return matrix[0][1]
	} else if matrix[1][2] == 1 {
		if matrix[1][0] < 0 {
			for i := 0; i < 3; i++ {
				matrix[1][i] -= (-1) * matrix[0][i]
			}
		}
		return matrix[1][0]
	} else {
		fmt.Print("Невозможно использовать данный ключ, подберите другой ")
		return -1
	}
}

// func printMatrix(matrix [2][3]int) {
// 	for _, row := range matrix {
// 		fmt.Printf("[%d, %d, %d]\n", row[0], row[1], row[2])
// 	}
// }
