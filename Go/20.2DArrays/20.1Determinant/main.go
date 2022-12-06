package main

import "fmt"

const rows = 3
const cols = 3

func determine(a [rows][cols]int) (result int) {
	d1 := a[0][0] * a[1][1] * a[2][2]
	d2 := a[0][0] * a[1][2] * a[2][1]
	d3 := a[0][1] * a[1][0] * a[2][2]
	d4 := a[0][1] * a[1][2] * a[2][0]
	d5 := a[0][2] * a[1][0] * a[2][1]
	d6 := a[0][2] * a[1][1] * a[2][0]
	result = d1 - d2 - d3 + d4 + d5 - d6
	fmt.Println(d1, d2, d3, d4, d5, d6)
	return
}

func main() {
	A := [rows][cols]int{
		{5, 2, 3},
		{4, 4, 7},
		{7, 1, 3},
	}
	determinant := determine(A)
	fmt.Println("Детерминант матрицы А равен", determinant)
}
