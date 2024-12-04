package main

import (
	"fmt"
)

// Filtragem com Ponteiros:Implemente uma função filtrarPares que recebe um ponteiro
//para um slice de inteiros e modifica o slice para conter apenas números pares.

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filtrarPares(&sl)
}

func filtrarPares(sl *[]int) {
	slice := *sl

	for i, v := range slice {
		if v%2 != 0 {
			slice = append(slice[:i], slice[i+1:]...)
			i--
		}
	}
	fmt.Println(slice)
}
