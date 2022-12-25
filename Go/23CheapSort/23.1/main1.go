package main

import "fmt"

const size = 10

func insertionSort(input [size]int) [size]int {
	for i := 1; i < size; i++ {
		j := i
		for j > 0 {
			if input[j-1] > input[j] {
				input[j-1], input[j] = input[j], input[j-1]
			}
			j--
		}

	}
	return input
}

func main() {
	fmt.Println("Задание 1. Сортировка вставками")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")

	unsorted := [size]int{10, 20, 5664, 646, 735, 89, 34, 67, 1, 8}
	sorted := insertionSort(unsorted)
	fmt.Println(sorted)

} /*Задание 1. Сортировка вставками
Что нужно сделать
Напишите функцию, сортирующую массив длины 10 вставками.*/
