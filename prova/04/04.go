package main

import (
	"fmt"
)

// Escreva uma função que remova elementos da posição 2 de uma slice de
// inteiros e retorne o novo slice.
func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println(remove(slice, 1))
}
func remove(slice []int, index int) []int {
	novoSlice := append(slice[:index], slice[index+1:]...)

	return novoSlice
}
