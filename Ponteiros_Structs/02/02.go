package main

import (
	"fmt"
)

//Swap de valores: cire uma função trocar que recebe dois ponteiros para int
//e troca os valores entre eles

func main() {
	a := 10
	b := 20

	pa := &a
	pb := &b

	swap(pa, pb)
}

func swap(pa, pb *int) {
	a := *pb
	b := *pa

	fmt.Println("\nValor de A:", a)
	fmt.Printf("Valor de B: %d\n", b)
}
