package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Конвейер")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")
	fmt.Println("Введите число или 'стоп' чтобы завершить программу")
	fmt.Println("первый коммит")
	var content string
	for {
		fmt.Scan(&content)
		if content == "стоп" {
			break
		}

		num, err := strconv.Atoi(content)
		if err != nil {
			fmt.Println("Ошибка:", err)
		}

		square := square(num)
		mult := multiplication(square)
		fmt.Println("Произведение:", <-mult)
	}
}

func square(num int) chan int {
	fChan := make(chan int)
	go func() {
		fChan <- num * num
	}()
	return fChan
}

func multiplication(fChan chan int) chan int {
	sChan := make(chan int)
	fChanRes := <-fChan
	fmt.Println("Квадрат:", fChanRes)
	go func() {
		sChan <- fChanRes * 2
	}()
	return sChan
}

/*Задание 1. Конвейер
Цели задания
Научиться работать с каналами и горутинами.
Понять, как должно происходить общение между потоками.
Что нужно сделать
Реализуйте паттерн-конвейер:

Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
Произведение: следующая горутина умножает квадрат числа на 2.
При вводе «стоп» выполнение программы останавливается.
Советы и рекомендации
Воспользуйтесь небуферизированными каналами и waitgroup.

Что оценивается
Ввод : 3

Квадрат : 9

Произведение : 18*/
