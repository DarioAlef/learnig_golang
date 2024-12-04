package main

import (
	"fmt"
)

//Inicialização de Structs: Implemente uma função novaPessoa que aloca dinamicamente
//uma struct Pessoa (com campos nome e idade) e retorna o ponteiro para a struct.

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	fmt.Println("\nDigite o seu nome: ")
	var nome string
	fmt.Scan(&nome)
	fmt.Println("\nDigite a sua idade: ")
	var idade int
	fmt.Scan(&idade)

	pessoa := novaPessoa(nome, idade)

	fmt.Println("\nNome: ", pessoa.Nome)
	fmt.Println("Idade: ", pessoa.Idade)
}

// INSTANCIAR
func novaPessoa(nome string, idade int) *Pessoa {
	return &Pessoa{
		Nome:  nome,
		Idade: idade,
	}
}
