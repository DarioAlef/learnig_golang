package main

import (
	"fmt"
)

// verifique se um vetor de 10 números está em ordem crescente
func main() {
	numbers := []int{}
	var num int

	fmt.Println("Digite um número (0 para parar): ")
	for {
		fmt.Scan(&num)
		if num == 0 {
			break
		}
		numbers = append(numbers, num)
	}

	//se o elemento do índice 'i' for maior que o elemento do índice seguinte, não está em ordem crescente
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] > numbers[i+1] {
			fmt.Println("O vetor nao esta em ordem crescente.")
			return
		}
	}

	fmt.Println("O vetor esta em ordem crescente.")
}
