package main

import (
	"fmt"
)

func anon(x int, y int, A func(x, y int) int) {
	defer fmt.Println(A(x, y))
}
func main() {
	anon(5, 2, func(x, y int) int { return x + y })
	anon(5, 2, func(x, y int) int { return x * y })
	anon(5, 2, func(x, y int) int { return x / y })
}
