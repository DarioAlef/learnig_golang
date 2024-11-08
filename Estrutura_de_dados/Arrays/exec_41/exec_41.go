package main

import (
	"fmt"
)

// leia um vetor de 15 números e retorne os que estão em posições ímpares
func main() {
	var num int
	numeros := make([]int, 0)

	for {
		fmt.Println("Digite um número (negativo para parar): ")
		fmt.Scan(&num)
		if num < 0 {
			break
		}
		numeros = append(numeros, num)
	}
	fmt.Println("Os números nas posições ímpares são:")
	for i, j := range numeros {
		if i%2 != 0 {
			fmt.Println(numeros[j])
		}
	}
}
