package main

import "fmt"

const size = 5

func main() {
	var x = [size]int{-10, -10, 55, 7, 50}
	var min, min2 int
	min = 4357378378453434
	min2 = 24673847689437
	for _, v := range x {
		if v < min {
			min2 = min
			min = v
		} else if min2 > v {
			min2 = v
		}
	}

	fmt.Println("Min:", min, "min2:", min2)
}
