package main

import "fmt"

//Manipulação de Slice por Ponteiro: escreva uma função addElement que recebe um ponteiro
//para slice de int e um valor int a ser adicionado. A função deve modificar o slice
//original adicionando um novo valor

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	add := 10

	addElement(&sl, add)
	fmt.Printf("\nSlice: %v\n", sl)
}

func addElement(p *[]int, add int) {
	*p = append(*p, add)
}
