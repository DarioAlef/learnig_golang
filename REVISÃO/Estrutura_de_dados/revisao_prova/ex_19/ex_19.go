package main

import (
	"fmt"
)

// Escreva uma função que multiplique duas matrizes 2x2. Faça com que a
// função retorne a matriz resultante da multiplicação
func main() {
	var mult [2][2]int
	matriz1 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	matriz2 := [2][2]int{
		{5, 6},
		{7, 8},
	}

	mult = multiplicacao(matriz1, matriz2)
	fmt.Printf("\nMatriz resultante da multiplicação: \n%v\n\t", mult)
}

func multiplicacao(matriz1 [2][2]int, matriz2 [2][2]int) [2][2]int {
	var mult [2][2]int

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz1); j++ {
			mult[i][j] = matriz1[i][j] * matriz2[i][j]
		}
	}
	return mult
}
