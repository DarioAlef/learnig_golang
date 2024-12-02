package main

import (
	"fmt"
)

// Faça uma função que receba uma array de inteiros e retorne a soma deles
func main() {
	arr := []int{}
	var aux1 int

	fmt.Println("\nDigite um número (negativo para parar): ")
	for {
		fmt.Scan(&aux1)
		if aux1 < 0 {
			break
		}
		arr = append(arr, aux1)
	}
	fmt.Println("\nO resultado é da soma desses números", somar(arr))
}

func somar(arr []int) int {
	var aux2 int
	for i := 0; i < len(arr)-1; i++ {
		aux2 = arr[i] + arr[i+1]
	}
	return aux2
}
