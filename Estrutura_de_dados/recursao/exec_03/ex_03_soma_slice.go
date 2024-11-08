package main

import "fmt"

//Escreva uma função recursiva que recebe um slice de
//inteiros e retorna a soma de todos os elementos.

func somar(s []int) int {
	if len(s) == 1 {
		return s[0]
	}
	return s[0] + somar(s[1:])
}
func main() {
	s := []int{1, 2, 3, 4, 5}
	fmt.Println(somar(s))
}
