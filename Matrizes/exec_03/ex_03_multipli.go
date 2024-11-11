package main

import (
	"fmt"
)

// multiplicação de matrizes
func main() {
	matriz1 := [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}}
	matriz2 := [3][3]int{
		{10, 11, 13},
		{12, 13, 14},
		{13, 14, 15},
	}
	multiplicacao(matriz1, matriz2)
	fmt.Println(multiplicacao(matriz1, matriz2))
}
func multiplicacao(matriz1 [3][3]int, matriz2 [3][3]int) [3][3]int {
	var mult [3][3]int

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2); j++ {
			mult[i][j] = matriz1[i][j] * matriz2[i][j]
		}
	}
	return mult
}
