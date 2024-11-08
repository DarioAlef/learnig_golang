package main

import (
	"fmt"
)

// inverta os elementos de um vetor de 10 números sem usar um vetor auxiliar
func main() {
	var n int
	numeros := [10]int{}

	for i := 0; i < 10; i++ {
		fmt.Printf("Digite um número: ")
		fmt.Scan(&n)
		numeros[i] = n
	}

	for i := 0; i < 10; i++ {
		fmt.Print(numeros[9-i], " , ")
	}
}

//O array numeros possui índices que vão de 0 a 9 (pois ele tem 10 elementos).
//9 - i é uma fórmula para acessar os elementos do array numeros de trás para frente:
//Quando i é 0: 9 - i resulta em 9, ou seja, numeros[9] (último elemento).
//Quando i é 1: 9 - i resulta em 8, ou seja, numeros[8].
//Quando i é 2: 9 - i resulta em 7, ou seja, numeros[7].
//E assim por diante, até que i seja 9, resultando em numeros[0] (primeiro elemento).
