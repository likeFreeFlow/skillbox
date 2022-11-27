package main

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	fmt.Println("Hello, World!")
	log.Info("tf")

}

/*Задача №1
Что нужно сделать
Выведите в консоль в формате json информацию о состоянии переменных `a` и `i` после вычисления значения переменной `a` на каждой итерации цикла в приведённом ниже примере.
func main() {
    for i := 0; i <= 10; i++ {
        a += i * i
    }
Пример содержимого файла с логами

{"level":"info","msg":"i = 0 a =  0","time":"2022-06-20T03:16:28+04:00"}

{"level":"info","msg":"i = 1 a =  1","time":"2022-06-20T03:16:28+04:00"}

{"level":"info","msg":"i = 2 a =  5","time":"2022-06-20T03:16:28+04:00"}

Советы и рекомендации
Воспользуйтесь пакетом Logru, для этого добавьте его в импорт.*/
