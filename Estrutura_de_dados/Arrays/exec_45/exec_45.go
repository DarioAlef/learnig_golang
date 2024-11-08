package main

import (
	"fmt"
)

// em um vetor de 8 elementos, multiplique todos os ímpares por 3
func main() {
	var n int
	numeros := [8]int{}

	for i := 0; i < 8; i++ {
		fmt.Printf("Digite um número: ")
		fmt.Scan(&n)
		numeros[i] = n
	}

	for i := 0; i < 8; i++ {
		if i%2 != 0 {
			numeros[i] = 3 * numeros[i]
		}
		fmt.Print(numeros[i], " , ")
	}

}
