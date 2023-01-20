package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Println("2. Graceful shutdown\n------------")

	chSign := make(chan os.Signal, 1)
	signal.Notify(chSign)
	chInt := make(chan int)

	go getN(chInt, chSign)
	for {
		select {
		case exit := <-chSign:
			log.Printf("Завершение программы (%v)", exit)
			goto pointClose
		default:
			n := <-chInt
			log.Printf("%v^2=%v\n", n, n*n)
		}
	}
pointClose:
	close(chSign)
	log.Println("Закрыт канал для сигналов")
	close(chInt)
	log.Println("Закрыт канал для чисел")

	log.Println("Программа завершена")
}

func getN(ch chan int, chSign chan os.Signal) {
	i := 0
	for {
		select {
		case exit := <-chSign:
			chSign <- exit
			log.Println("Завершена функция получения чисел")
			return
		default:
			i += 1
			ch <- i
		}
	}
}
