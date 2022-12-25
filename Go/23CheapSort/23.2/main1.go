package main

import "fmt"

const size = 10

func main() {
	fmt.Println("Задание 2. Анонимные функции")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")

	bubbleSort := func(input [size]int) [size]int {
		for i := size; i > 0; i-- {
			for j := 1; j < i; j++ {
				if input[j-1] < input[j] {
					input[j-1], input[j] = input[j], input[j-1]
				}
			}
		}
		return input
	}
	unSorted := [size]int{10, 20, 5664, 646, 735, 89, 34, 67, 1, 8}
	sorted := bubbleSort(unSorted)
	fmt.Println(sorted)

}

/*
Задание 2. Анонимные функции
Что нужно сделать
Напишите анонимную функцию, которая на вход получает массив типа integer,
сортирует его пузырьком и переворачивает
(либо сразу сортирует в обратном порядке, как посчитаете нужным).*/
