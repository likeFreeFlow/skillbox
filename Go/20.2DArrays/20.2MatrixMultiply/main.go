package main

import "fmt"

func multiply(a [3][5]int, b [5][4]int) (c [3][4]int) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b[0]); j++ {
			for k := 0; k < len(a[0]); k++ {
				c[i][j] += a[i][k] * b[k][j]
				fmt.Println(i, j, k)
			}
			fmt.Println(c[i][j])
		}
	}
	return
}
func main() {
	fmt.Println("Задание 2. Умножение матриц\n")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru \n\n")

	a := [3][5]int{
		{0, 1, 2, 3, 4},
		{5, 6, 7, 8, 9},
		{10, 11, 12, 13, 14},
	}
	b := [5][4]int{
		{100, 101, 102, 103},
		{104, 105, 106, 107},
		{108, 109, 110, 111},
		{112, 113, 114, 115},
		{116, 117, 118, 119},
	}
	c := multiply(a, b)
	fmt.Println(c)

}

/*Задание 2. Умножение матриц
Что нужно сделать
Напишите функцию, умножающую две матрицы размерами 3 × 5 и 5 × 4.

Советы и рекомендации
Алгоритм умножения матриц проиллюстрирован в «Википедии».
В качестве среды разработки может помочь Goland или VScode.
Что оценивается
Правильность размеров исходных и конечной матрицы.
Арифметическая правильность ответа.*/