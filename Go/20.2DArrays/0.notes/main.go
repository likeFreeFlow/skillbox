package main
import "fmt"

const rows=3
const size=4
func maximum (input [rows][size]int) int{
	max:=input[0]
	for i:=0;i<len(input);i++{
		//len возвращает rows, 1 параметр 2д массива
		//len(matrix[i]) возвращает size
		if input[i]>max{
			max=input[i]
		}
		return max
	}

}
func main () {
	matrix:=[rows][size]
}