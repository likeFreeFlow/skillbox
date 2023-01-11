package main

import (
	"fmt"
	"strings"
)

func searchForIndex(whereToLook [4]string, whatToLook [5]rune) (index [4][5]int) {
	for i := 0; i < len(whereToLook); i++ {
		words := strings.Fields(whereToLook[i])
		word := words[1]
		for j := 0; j < len(whatToLook); j++ {
			//strings.index-узнал после того, как расписал свое решение
			index[i][j] = -1
			for k := 0; k < len(word); k++ {
				if whatToLook[j] == rune(word[k]) {
					index[i][j] = k
					break
				}

			}
		}
	}
	return
}
func printer(toPrint [4][5]int, toCompare [5]rune, original [4]string) {
	for i := 0; i < len(original); i++ {
		fmt.Println("Строка поиска", original[i])
		for j := 0; j < len(toCompare); j++ {
			fmt.Print("Индекс вхождения символа ")
			fmt.Printf(" %s  %d\n", string(toCompare[j]), toPrint[i][j])
		}
	}
}
func main() {
	fmt.Println("Задание 2. Поиск символов в нескольких строках")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")

	whereToLook := [4]string{"Hello world", "Hello Skillbox",
		"Привет Мир", "Привет Skillbox"}

	whatToLook := [5]rune{'w', 'b', 'L', 'р', 'x'}
	index := searchForIndex(whereToLook, whatToLook)
	printer(index, whatToLook, whereToLook)

}

/*Задание 2. Поиск символов в нескольких строках
Что нужно сделать
Напишите функцию, которая на вход принимает массив предложений
 (длинных строк) и массив символов типа rune, а возвращает 2D-массив,
 где на позиции [i][j] стоит индекс вхождения символа j из chars в последнее
  слово в предложении i (строку надо разбить на слова и взять последнее).
  То есть сигнатура следующая:

func parseTest(sentences []string, chars []rune)

Советы и рекомендации*/
