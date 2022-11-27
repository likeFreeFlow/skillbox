package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Задание 1. Подсчёт чётных и нечётных чисел в массиве\n")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru \n\n")
	rand.Seed(time.Now().UnixNano())
	var tenRandoms [10]int
	for i := 0; i < 10; i++ {
		tenRandoms[i] = rand.Int() //
	}
	var cntEven, cntUnEven int
	for _, v := range tenRandoms {
		if v%2 == 0 {
			cntEven++
		} else {
			cntUnEven++
		}
	}
	fmt.Printf("Чётных:%d \nНечетных:%d", cntEven, cntUnEven)
}

/*Задание 1. Подсчёт чётных и нечётных чисел в массиве
Что нужно сделать
Разработайте программу, позволяющую ввести 10 целых чисел, а затем вывести из них количество чётных и нечётных чисел. Для ввода и подсчёта используйте разные циклы.

Что оценивается
Для введённых чисел 1, 1, 1, 2, 2, 2, 3, 3, 3, 4 программа должна вывести: чётных — 4, нечётных — 6.*/
