package main

import (
	"fmt"
	"math/rand"
	"time"
)

func createArray() (output [6]int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 6; i++ {
		output[i] = rand.Intn(100)

	}
	return
}
func bubbleIt(input [6]int) [6]int {
	for i := 1; i < 6; i++ {
		if input[i-1] >= input[i] {
			input[i-1], input[i] = input[i], input[i-1]
			i = 0
		}
	}
	return input
}
func showMeTheArray(input [6]int) {
	for _, v := range input {
		fmt.Println(v)
	}
}
func main() {
	fmt.Println("Задание 1. Слияние отсортированных массивов\n")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru \n\n")
	arrayToSort := createArray()
	bubbledOne := bubbleIt(arrayToSort)
	showMeTheArray(bubbledOne)
}
