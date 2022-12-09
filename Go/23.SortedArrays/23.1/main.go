package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func fillMeUp() (a [size]int) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range a {
		a[i] = rand.Int()
	}
	return
}

func main() {
	fmt.Println("Задание 1. Расчёт по формуле")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")
	lookForMe := 5
	checkMe := fillMeUp()
	checkMe[0] = 5
	//вместо вывода массива "для наглядности" можно искусственно задать местоположение нашего искомого элемента
	found := false
	for i, v := range checkMe {
		if v == lookForMe {
			fmt.Println("Число найдено. В массиве после введенного числа находится еще", size-i-1, "чисел")
			found = true
			break
		}

	}
	if !found {
		fmt.Println("0")
	}

}

/*Задание 1. Подсчёт чисел в массиве
Что нужно сделать
Заполните массив неупорядоченными числами на основе генератора случайных чисел. Введите число. Программа должна найти это число в массиве и вывести, сколько чисел находится в массиве после введённого. При отсутствии введённого числа в массиве — вывести 0. Для удобства проверки реализуйте вывод массива на экран.

Что оценивается
Программа должна корректно считать числа в массиве после заданного.*/
