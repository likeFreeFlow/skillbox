package main

import (
	"flag"
	"fmt"
)

var Found bool = false

func lookForSubString(str string, sstr string) {

	sRune := []rune(str)
	ssRune := []rune(sstr)
	for i, v := range sRune {
		if v == ssRune[0] {
			for j := 1; j < len(ssRune); j++ {
				if sRune[i+j] == ssRune[j] {
					if j == len(ssRune)-1 {
						Found = true
						break
					}
				}
			}

		}

	}
	fmt.Println(Found)
}

func main() {
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")
	var whereToLook string
	var whatToLook string

	flag.StringVar(&whereToLook, "str", "", "set string")
	flag.StringVar(&whatToLook, "substr", "", "set substring")
	flag.Parse()
	lookForSubString(whereToLook, whatToLook)

}

/*25.9 Практическая работа
Цель задания
Написать программу для нахождения подстроки в кириллической подстроке.
Программа должна запускаться с помощью команды:

 go run main.go --str "строка для поиска" --substr "поиска"

Для реализации такой работы с флагами воспользуйтесь пакетом flags,
а для поиска подстроки в строке вам понадобятся руны.

Что нужно сделать
Спроектировать алгоритм поиска подстроки.
Определить строку и подстроку, используя флаги.
Написать алгоритм реализацию для работы со строками UTF-8
(для этого необходимо воспользоваться рунами).*/
