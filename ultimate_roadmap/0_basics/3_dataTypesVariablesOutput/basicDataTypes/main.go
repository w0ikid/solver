package main

import (
	"fmt"
)

func main() {
	var i int
	var ll int64
	var ch byte
	var f float32
	var d float64

	// Считывание всех значений с одной строки
	fmt.Scanf("%d %d %c %f %f", &i, &ll, &ch, &f, &d)

	// Вывод значений на отдельных строках
	fmt.Println(i)
	fmt.Println(ll)
	fmt.Println(string(ch)) // Преобразуем byte в строку для корректного вывода
	fmt.Printf("%.2f\n", f) // Округляем float32 до двух знаков после запятой
	fmt.Printf("%.1f\n", d) // Округляем float64 до одного знака после запятой
}
