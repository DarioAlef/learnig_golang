package main

import (
	"fmt"
)

//Ponteiros e Strings: faça uma função reverterString que receba um ponteiro para uma
//string e altere o conteúdo para a string invertida

func main() {
	fmt.Println("\nDigite uma palavra: ")
	var palavra string
	fmt.Scan(&palavra)

	reverterString(&palavra)
}

func reverterString(palavra *string) {
	w := *palavra
	sl := make([]string, 0)

	for i := len(w) - 1; i >= 0; i-- {
		sl = append(sl, string(w[i]))
	}

	fmt.Println(sl)
}
