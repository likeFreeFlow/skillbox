package main
import "fmt"

const rows=3
const size=4
func printOneCycle (m[rows][cols]int){
	//пробежка в 1 цикл
	for i:=0;i<rows*cols;i++{
		row:=i/cols
		col:=i%cols
	}

}
func transpose(A[rows][cols]int)transposed [cols][rows]int{
	for i:=0;i<len(a[0]);i++{
		for j:=0;j<len(A);j++{
			transposed[i][j]=A[j][i]
		}
	}
}
func diagonalSum(A[rows][cols]int)sum int{
	for i:=0;i,len(A);i++{
		sum+=a[i][i]
	}
}
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