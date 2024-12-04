package main

import (
	"fmt"
)

//Map com Ponteiro: Crie uma função adicionarRegistro que receba um ponteiro para
//um map[string]int, uma chave e um valor, e adicione o par chave-valor ao mapa

func main() {
	pessoa := make(map[string]int)

	adicionarRegistro(&pessoa, "José", 37)
}

func adicionarRegistro(pessoa *map[string]int, chave string, valor int) {
	(*pessoa)[chave] = valor

	fmt.Println(pessoa)
}
