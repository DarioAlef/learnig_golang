package main

import (
	"fmt"
)

//Alocação Dinâmica e Tamanho Variável: faça uma função alocarMatriz que aloca uma matriz
//dinâmica (slice de slices) de n por m e retorna um ponteiro para a matriz

func main() {
	fmt.Println("\nDigite o número de linhas: ")
	var l int
	fmt.Scan(&l)
	fmt.Println("Digite o número de colunas: ")
	var c int
	fmt.Scan(&c)

	(alocarMatriz(l, c))

	fmt.Println("\nMatriz")
	for _, v := range *alocarMatriz(l, c) {
		fmt.Println(v)
	}
}

func alocarMatriz(l, c int) *[][]int {
	matriz := make([][]int, l)

	for i := 0; i < l; i++ {
		matriz[i] = make([]int, c)
	}

	for i := 0; i < l; i++ {
		for j := 0; j < c; j++ {
			matriz[i][j] = i + 1
			matriz[i][j] = j + 1
			matriz[i][j] = i + j
		}
	}

	return &matriz
}
