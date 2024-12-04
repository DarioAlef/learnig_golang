package main

import (
	"fmt"
)

//Cálculo Recursivo com Ponteiros:faça uma função recursiva somarAteN que recebe um ponteiro
//para um int e retorna a soma de todos os números de 1 até o valor apontado pelo ponteiro

func main() {
	fmt.Println("\nDigite um número: ")
	var n int
	fmt.Scan(&n)

	fmt.Println(somarAteN(&n))
}

func somarAteN(n *int) int {
	aux := *n
	*n = *n - 1

	if aux <= 0 {
		return 0
	} else {
		return aux + somarAteN(n)
	}
}
