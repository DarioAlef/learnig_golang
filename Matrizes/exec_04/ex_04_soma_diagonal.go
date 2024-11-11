package main

import "fmt"

// implemente uma função que recebe uma matriz 3x3 e
// retorna a soma dos elementos da diagonal principal
func main() {
	var matriz [3][3]int

	fmt.Println("Digite os 9 valores da matriz 3x3: ")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("Digite o valor para a posição [%d][%d] ->", i, j)
			fmt.Scan(&matriz[i][j])
		}
	}
	somadp(matriz)
	fmt.Println("\nA soma dos elementos da matriz é:", somadp(matriz))
}
func somadp(matriz [3][3]int) int {
	var dp int
	var n int
	for n = 0; n < 3; n++ {
		dp += matriz[n][n]
	}
	return dp
}
