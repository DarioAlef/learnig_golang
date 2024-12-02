package main

import (
	"fmt"
)

//Soma de Elementos por Referência: faça uma função somaSlice que recebe um ponteiro
//para um slice de int e retorna a soma de todos os elementos

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(somaSlice(&slice))
}

func somaSlice(slice *[]int) int {
	sl := *slice
	var soma int

	for _, v := range sl {
		soma = soma + v
	}
	return soma
}
