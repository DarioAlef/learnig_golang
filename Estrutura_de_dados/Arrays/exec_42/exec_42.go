package main

import (
	"fmt"
)

// mova todos os números negativos de um vetor de 10 elementos para o final
func main() {
	var numeros [10]int

	// Lendo os números
	fmt.Println("Digite 10 números inteiros:")
	for i := 0; i < 10; i++ {
		fmt.Scan(&numeros[i])
	}

	//len(numeros)-1 lê de trás pra frente
	for i := 0; i < len(numeros)-1; i++ {
		minimo_index := i
		//':= i + 1' para coincidir com os índices iniciais
		for j := i + 1; j < len(numeros); j++ {
			if numeros[j] < numeros[minimo_index] {
				minimo_index = j
			}
		}
		numeros[i], numeros[minimo_index] = numeros[minimo_index], numeros[i]
	}

	// Imprimindo os números ordenados
	fmt.Printf("Números ordenados: %d", numeros)
}
