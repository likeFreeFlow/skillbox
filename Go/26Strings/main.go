package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Цель задания - Написать программу аналог cat. ")
	fmt.Println("Ищу практику в интересных проектах на го. barakovvm@ya.ru ")
	flag.Parse()
	fileNames := flag.Args()
	if len(fileNames) == 1 {
		f1, err := os.Open(fileNames[0])
		if err != nil {
			panic(err)
		}
		defer f1.Close()
		scanner := bufio.NewScanner(f1)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	if len(fileNames) == 2 {
		f1, err := os.Open(fileNames[0])
		if err != nil {
			panic(err)
		}
		defer f1.Close()
		scanner := bufio.NewScanner(f1)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		fmt.Println()
		f2, err := os.Open(fileNames[1])
		if err != nil {
			panic(err)
		}
		defer f2.Close()
		scanner = bufio.NewScanner(f2)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}
	if len(fileNames) == 3 {
		f3, err := os.Create(fileNames[2])

		if err != nil {
			fmt.Println("Unable to create file:", err)
			os.Exit(1)
		}
		defer f3.Close()

		f1, err := os.Open(fileNames[0])
		if err != nil {
			panic(err)
		}
		defer f1.Close()
		scanner := bufio.NewScanner(f1)
		for scanner.Scan() {
			f3.WriteString(scanner.Text())
			f3.WriteString(("\n"))
		}
		f3.WriteString(("\n"))

		f2, err := os.Open(fileNames[1])
		if err != nil {
			panic(err)
		}
		defer f2.Close()
		scanner = bufio.NewScanner(f2)
		for scanner.Scan() {
			f3.WriteString(scanner.Text())
			f3.WriteString(("\n"))
		}
		fmt.Println("Done.")
	}
}

/*Цель задания
Написать программу аналог cat.
Что нужно сделать
Программа должна получать на вход имена двух файлов,
необходимо  конкатенировать их содержимое, используя strings.Join.

Что оценивается
При получении одного файла на входе программа должна печатать его содержимое на экран.
При получении двух файлов на входе программа соединяет их и печатает содержимое обоих
файлов на экран.
Если программа запущена командой go run firstFile.txt secondFile.txt resultFile.txt,
то она должна написать два соединённых файла в результирующий.*/
