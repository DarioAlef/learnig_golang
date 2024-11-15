package main

import (
	"fmt"
)

// Crie uma função recursiva que calcule a potência de um número
// (base e expoente). Por exemplo, 2^3 deve retornar 8
func main() {
	var base int
	var expo int

	fmt.Println("\nDigite a base: ")
	fmt.Scan(&base)
	fmt.Println("Digite o expoente: ")
	fmt.Scan(&expo)

	fmt.Printf("O resultado da potência é: %d\n", pot(base, expo))
}

func pot(base int, expo int) int {
	if expo == 0 {
		return 1
	}

	if expo == 1 {
		return base
	}

	//base^expo = base^(expo/2) * base^(expo/2)
	//temos := base^expo = pot(base, expo/2) * pot(base, expo/2)
	termo := pot(base, expo/2)
	if expo%2 == 0 {
		return termo * termo
	} else {
		return base * termo * termo
	}
}
