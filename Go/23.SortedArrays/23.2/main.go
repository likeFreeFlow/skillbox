package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

const size = 12

func fillMeUp() (a [size]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		a[i] = i + rand.Intn(2)
	}
	return
}

func findIt(x int, a [size]int) {
	i := sort.Search(len(a), func(i int) bool { return a[i] >= x })
	if i < len(a) && a[i] == x {
		fmt.Printf("found %d at index %d in %v\n", x, i, a)
	} else {
		fmt.Printf("%d not found in %v\n", x, a)
	}
}

func main() {
	fmt.Println("Задание 2. Нахождение первого вхождения числа в упорядоченном массиве")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")
	lookForMe := 2
	whereToLook := fillMeUp()
	fmt.Println(whereToLook)
	fmt.Scan(&lookForMe)
	findIt(lookForMe, whereToLook)
	fmt.Println(whereToLook)
}

/*Задание 2. Нахождение первого вхождения числа в упорядоченном массиве (числа могут повторяться)
Что нужно сделать
Заполните упорядоченный массив из 12 элементов и введите число. Необходимо реализовать поиск первого вхождения заданного числа в массив. Сложность алгоритма должна быть минимальная.

Что оценивается
Верность индекса.

При вводе массива 1 2 2 2 3 4 5 6 7 8 9 10 и вводе числа 2 программа должна вывести индекс 1.*/
