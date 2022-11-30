package main

import (
	"fmt"
	"math/rand"
	"time"
)

//const arrayLen = 5

func mergeArrays(a [5]int, b [4]int) (output [9]int) {

	for i := 0; i < 4; i++ {
		output[i*2] = a[i]
		output[i*2+1] = b[i]
	}
	output[8] = a[4]
	return
}

/*func createArray() (a [arrayLen]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < arrayLen+1; i++ {
		a[i] = rand.Int()
	}
	return a
}

func bubbleIt(input [arrayLen]int) [arrayLen]int {
	for i := 1; i < arrayLen; i++ {
		if input[i-1] > input[1] {
			input[i-1], input[i] = input[i], input[i]-1
			i--
		}
	}
	return input
}*/
func main() {
	fmt.Println("Задание 1. Слияние отсортированных массивов\n")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru \n\n")
	rand.Seed(time.Now().UnixNano())
	//не получилось создать массив внутри функции, тк переменная не принимается уместной
	//для создания массива[с переменной в размере]
	//я не понимаю, как решить эту проблему и пришлось написать все в теле программы
	//буду благодарен за помощь в решении вопроса
	//переделал все по простому, убил кучу времени
	//бесит
	/*a := createArray
	b := createArray
	a = bubbleIt(a)
	b= bubbleIt(b)
	c := mergeArrays(a, b)*/
	var a = [5]int{1, 3, 5, 7, 9}
	var b = [4]int{2, 4, 6, 7}
	c := mergeArrays(a, b)
	fmt.Println(c)

}
