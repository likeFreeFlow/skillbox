package main

import "fmt"

const size = 12

func sortIt(a [size]int) (even [size]int, unEven [size]int) {
	var eC, ueC int
	for _, v := range a {
		if v%2 == 0 {
			even[eC] = v
			eC++
		} else {
			unEven[ueC] = v
			ueC++
		}
	}
	return
}
func main() {
	var unsorted = [size]int{5, 4, 9, 2, 4, 6, 7, 9, 0, 3, 2, 3}
	even, unEven := sortIt(unsorted)
	fmt.Println("Массив четных чисел", even)
	fmt.Println("Массив нечетных чисел", unEven)

}

/*Задание 1. Чётные и нечётные
Что нужно сделать
Напишите функцию, которая принимает массив чисел,
 а возвращает два массива: один из чётных чисел, второй из нечётных. */
