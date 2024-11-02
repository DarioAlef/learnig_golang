package main

import (
	"fmt"
	"strings"
)

// verifique se uma palavra inserida pelo usuário é um palíndromo.
func main() {
	var palavra string
	var reverso string

	fmt.Println("Digite um palavra qualquer: ")
	fmt.Scan(&palavra)

	palavra = strings.ToLower(strings.ReplaceAll(palavra, " ", ""))

	//for i := 0; i < len(palavra)/2; i++ {
	//	if palavra[i] != palavra[len(palavra)-i-1] {
	//	}
	//}
	for i := len(palavra) - 1; i >= 0; i-- {
		reverso += string(palavra[i])
		//i começa pelo tamanho de palavra (-1) para ser ao contrário
		//adicionando à variável array 'reverso'
	}
	if reverso == palavra {
		fmt.Println("É um palíndronomo")
	} else {
		fmt.Println("Não é um palíndromo")
	}
}
