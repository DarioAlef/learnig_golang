package main

import (
	"fmt"
)

//Soma Recursiva com Ponteiro para Slice: faça uma função recursiva somaRecursiva que recebe
//um ponteiro para um slice de int e calcula a soma dos elementos

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	somaRecursiva(&slice)
}

func somaRecursiva(slice *[]int) {
	sl := *slice
	var soma int

	for _, v := range sl {
		soma = soma + v
	}
	fmt.Println(soma)
}
