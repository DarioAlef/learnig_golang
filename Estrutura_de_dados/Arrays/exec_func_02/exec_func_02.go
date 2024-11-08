package main

import (
	"fmt"
)

// faça uma função que aceite um array de inteiros e retorne o valor máximo
func main() {
	arr := []int{}
	var aux int

	fmt.Println("\nDigite um número (negativo para parar): ")
	for {
		fmt.Scan(&aux)
		if aux < 0 {
			break
		}
		arr = append(arr, aux)
	}
	fmt.Println("\nO maior número é", max(arr))
}
func max(arr []int) int {
	var aux2 int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			aux2 = arr[i]
		} else {
			aux2 = arr[i+1]
		}
	}
	return aux2
}
