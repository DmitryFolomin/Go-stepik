package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	slice := make([]float64, N)

	// Считываем элементы массива
	for i := 0; i < N; i++ {
		fmt.Scan(&slice[i])
	}

	// Вычисляем сумму элементов
	var sum float64
	for _, value := range slice {
		sum += value
	}

	// Вычисляем среднее арифметическое
	average := sum / float64(N)

	// Выводим результат
	fmt.Printf("%.2f\n", average)
}
