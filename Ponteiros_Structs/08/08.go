package main

import (
	"fmt"
)

//Recursão com Ponteiros: escreva uma função recursiva fatorialPtr que recebe um ponteiro
//para um int e calcula o fatorial do valor apontado, substituindo o valor pelo resultado

func main() {
	var n int

	fmt.Println("Digite um número inteiro qualquer: ")
	fmt.Scan(&n)
	fmt.Println(fatorialPtr(&n))
}

func fatorialPtr(ptr *int) int {
	aux := *ptr

	*ptr = *ptr - 1

	if aux == 0 {
		return 1
	} else {
		return aux * fatorialPtr(ptr)
	}
}
