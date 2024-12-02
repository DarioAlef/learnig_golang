package main

import (
	"fmt"
)

// Faça uma função que conte quantos números pares há em um slice.
func main() {
	var numbers []int
	var num int

	fmt.Println("\nDigite um número ('0' para parar): ")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numbers = append(numbers, num)
	}

	count := pares(numbers)
	fmt.Println("Quantidade de números pares: ", count)
}

func pares(numbers []int) int {
	var count int
	for _, n := range numbers {
		if n%2 == 0 {
			count++
		}
	}
	return count
}
