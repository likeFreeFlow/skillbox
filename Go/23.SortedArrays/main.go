package main

import (
	"fmt"
	"math/rand"
	"time"
)

const size = 10

func fillMeUp() (a [size]int) {
	rand.Seed(time.Now().UnixNano())
	for i, _ := range a {
		a[i] = rand.Int()
	}
	return
}

func main() {
	lookForMe := 5
	checkMe := fillMeUp()
	checkMe[0] = 5
	//вместо вывода массива можно искусственно задать местоположение нашего искомого элемента
	found := false
	for i, v := range checkMe {
		if v == lookForMe {
			fmt.Println("Число найдено. В массиве после введенного числа находится еще", size-i-1, "чисел")
			found = true
			break
		}

	}
	if !found {
		fmt.Println("0")
	}

}
