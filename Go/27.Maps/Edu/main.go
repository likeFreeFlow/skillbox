package main

import (
	"fmt"
	"person/person"
	"strconv"
)

/*//Пример 1. Объявление метода
func (p Person) Print() {
	fmt.Println(p)
}

//Пример 1. Объявление функции
func PrintInfo(p Person) {
	fmt.Println(p)
}*/

/*//Пример 2. Печать структуры методом
//через значение:
func (p.Person) SetNameByValue(n string) {
	p.Name=n
}
//Через указатель:
func (p.*Person) SetNameByPointer(n string) {
	p.Name=n
}*/
type (
	MyInt int
)

func (mi MyInt) toString() string {
	return (strconv.Itoa(int(mi)))

}
func main() {
	p := person.Person{
		Age:  38,
		Name: "Corado",
	}
	//печать через функцию
	//PrintInfo(p)
	//печать через метод
	//p.Print

	p.SetNameByValue("Tony")
	fmt.Println(p)
	p.SetNameByPointer("Tony")
	fmt.Println(p)
	var i MyInt = 10
	iStr := i.toString()
	fmt.Println(iStr)
}
