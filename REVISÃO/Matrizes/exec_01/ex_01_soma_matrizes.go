package main

import "fmt"

// escreva uma função que recebe duas matrizes 3x3 como parâmetros e retorna uma nova
// matriz 3x3 contendo a soma das duas matrizes
func main() {
	matriz1 := [3][3]int{
		{1, 1, 1},
		{1, 1, 1},
		{1, 1, 1},
	}
	matriz2 := [3][3]int{
		{2, 2, 2},
		{2, 2, 2},
		{2, 2, 2},
	}
	//soma = matrix1[0][0] + matrix2[0][0]

	sum(matriz1, matriz2)
	fmt.Println(sum(matriz1, matriz2))
}

func sum(matriz1 [3][3]int, matriz2 [3][3]int) [3][3]int {
	var soma [3][3]int

	for i := 0; i < len(matriz1); i++ {
		for j := 0; j < len(matriz2[i]); j++ {
			soma[i][j] = matriz1[i][j] + matriz2[i][j]
		}
	}
	return soma
}

//// Semente baseada no tempo atual
//rand.Seed(time.Now().UnixNano())
//
//// Sorteia um número entre 0 e 100
//ale := rand.Intn(101)
