package main

import (
	"fmt"
	"math/rand"
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
func main() {
	fmt.Println("Задание 1. Расчёт по формуле")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")
	//var sortedNotUnique=[...]int{1, 2, 2, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	lookForMe := 2
	whereToLook := fillMeUp()
	for i, v := range whereToLook {
		if v == lookForMe {
			fmt.Println("Индекс элемента ", i)
			break
		}
	}
	fmt.Println(whereToLook)
}
