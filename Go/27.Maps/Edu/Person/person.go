package person

type Person struct {
	Age  int
	Name string
}

//Пример 2. Печать структуры методом
//через значение:
func (p Person) SetNameByValue(n string) {
	p.Name = n
}

//Через указатель:
func (p *Person) SetNameByPointer(n string) {
	p.Name = n
}
