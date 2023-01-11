package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type StudentType struct {
	name string

	age int

	grade int
}

//метод инфо на возврат информации о студенте
func (s StudentType) info() string {
	return s.name + " " + strconv.Itoa(s.age) + " " + strconv.Itoa(s.grade)
}

//Переделать на методы
//методы put&get с проверкой на ошибки

func (ivanov *StudentType) put(m map[string]*StudentType) {
	m[ivanov.name] = ivanov
}

func (ivanov StudentType) get(m map[string]*StudentType) bool {
	_, out := m[ivanov.name]
	return out
}

func main() {

	mapStudents := make(map[string]*StudentType, 0)

	// Присваивание значений переменным для посл проверки и занесения в карту
	fmt.Println("Введите имя студента, возраст и курс через пробел")
	for {

		input := bufio.NewReader(os.Stdin)
		line, err := input.ReadString('\n')
		if err == io.EOF {
			fmt.Print("-\n")
			fmt.Println("---------------------")
			fmt.Println("Ввод окончен! Список студентов:")
			break
		}
		lineSplitted := strings.Fields(line)
		studentName := lineSplitted[0]
		if lineSplitted[0] == "eof" {
			fmt.Print("-\n")
			fmt.Println("---------------------")
			fmt.Println("Ввод окончен! Список студентов:")
			break
		}
		studentAge, errAge := strconv.Atoi(lineSplitted[1])
		studentGrade, errGrade := strconv.Atoi(lineSplitted[2])
		if errAge != nil || errGrade != nil {
			fmt.Print("У студента проблемы с возрастом и курсом. А у тебя с головой. Введи еще раз и правильно\n")
			continue
		}
		ivanov := StudentType{
			name:  studentName,
			age:   studentAge,
			grade: studentGrade,
		}
		if err := ivanov.get(mapStudents); err == false {
			ivanov.put(mapStudents)
		} else {
			fmt.Print("А новые студенты есть? Я дублей не потерплю!\n")
		}
	}
	counter := 1
	for _, v := range mapStudents {
		fmt.Printf("%v. %s\n", counter, v.info())
		counter++
	}
	//fmt.Println(line /*studentName, studentAge, studentGrade*/)
	/*for {
		fmt.Println("Введите имя студента")
		studentName := bufio.NewReader(os.Stdin)
		//fmt.Println("Возраст")
		//studentAge := bufio.NewReader(os.Stdin)
		//fmt.Println("Курс")
		//studentGrade := bufio.NewReader(os.Stdin)

		fmt.Println(studentName/*, studentAge, studentGrade)
		break
	}*/

	//fmt.Println(Student)
}

/*Цель задания
Научиться работать с композитными типами данных: структурами и картами

Что нужно сделать
Напишите программу, которая считывает ввод с stdin, создаёт структуру student и
записывает указатель на структуру в хранилище map[studentName] *Student.

Программа должна получать строки в бесконечном цикле, создать структуру Student
через функцию newStudent, далее сохранить указатель на эту структуру в map,
а после получения EOF (ctrl + d) вывести на экран имена всех студентов из хранилища.
Также необходимо реализовать методы put, get. */
