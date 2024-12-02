package main

import (
	"fmt"
	"sort"
)

// Crie uma função para ordenar um slice de inteiros em ordem crescente.
// Utilize o pacote sort de Go
func main() {
	var numeros []int
	var n int

	fmt.Println("Digite um número ('0' para parar): ")
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}
		numeros = append(numeros, n)
	}
	fmt.Println("Os números em ordem crescente: ", ordenar(numeros))
}

func ordenar(numeros []int) []int {
	sort.Ints(numeros)
	return numeros
}
