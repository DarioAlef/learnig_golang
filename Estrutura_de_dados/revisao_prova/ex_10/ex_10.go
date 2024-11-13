package main

import (
	"fmt"
)

// Faça uma função que conte quantos números pares há em um slice.
func main() {
	numbers := []int{}
	var num int

	fmt.Println("Digite uma sequência de inteiros (0 para parar): ")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numbers = append(numbers, num)
	}
	pares(numbers)
}
func pares(numbers []int) []int {
	nPares := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			nPares = append(nPares, num)
		}
	}
	return nPares
}
