package main

import "fmt"

// leia uma lista de números inteiros e os ordene em ordem crescente.
func main() {
	var numeros []int

	fmt.Println("Digite os números (digite um número negativo para parar):")
	for {
		var num int
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		numeros = append(numeros, num)
	}

	// Bubble Sort
	for i := 0; i < len(numeros)-1; i++ {
		for j := 0; j < len(numeros)-i-1; j++ { //compara ao contrário
			if numeros[j] > numeros[j+1] {
				// Troca os elementos de lugar
				numeros[j], numeros[j+1] = numeros[j+1], numeros[j]
			}
		}
	}

	fmt.Println("Números ordenados:", numeros)
}
